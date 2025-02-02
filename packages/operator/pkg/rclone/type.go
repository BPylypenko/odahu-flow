//
//    Copyright 2019 EPAM Systems
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package rclone

import (
	"context"
	"fmt"
	odahuflowv1alpha1 "github.com/odahu/odahu-flow/packages/operator/api/v1alpha1"
	"github.com/odahu/odahu-flow/packages/operator/pkg/apis/connection"
	"github.com/pkg/errors"
	_ "github.com/rclone/rclone/backend/local" // local specific handlers
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/cache"
	"github.com/rclone/rclone/fs/operations"
	"github.com/rclone/rclone/fs/sync"
	uuid "github.com/satori/go.uuid"
	"path"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

const (
	copyEmptySrcDirs = true
)

var log = logf.Log.WithName("rclone")

type ObjectStorage struct {
	RemoteConfig *FileDescription
}

type FileDescription struct {
	FsName string
	Path   string
}

func NewObjectStorage(conn *odahuflowv1alpha1.ConnectionSpec) (obj *ObjectStorage, err error) {
	name := uuid.NewV4()

	return NewObjectStorageWithName(name.String(), conn)
}

func NewObjectStorageWithName(name string, conn *odahuflowv1alpha1.ConnectionSpec) (obj *ObjectStorage, err error) {
	var config *FileDescription

	switch conn.Type {
	case connection.S3Type:
		config, err = createS3config(name, conn)
	case connection.GcsType:
		config, err = createGcsConfig(name, conn)
	case connection.AzureBlobType:
		config, err = createAzureBlobConfig(name, conn)
	default:
		return nil, errors.New(fmt.Sprintf("Unexpected connection type: %s", conn.Type))
	}

	log.Info("Extract FileDescription", "file description", config)

	if err != nil {
		return nil, err
	}

	return &ObjectStorage{RemoteConfig: config}, nil
}

// newFsFile creates a Fs from a name but may point to a file.
// It returns a string with the file name if points to a file
// otherwise "".
func newFsFile(remote string) (fs.Fs, string, error) {
	_, _, fsPath, err := fs.ParseRemote(remote)
	if err != nil {
		cntRrr := fs.CountError(err)
		if cntRrr != nil {
			return nil, "", cntRrr
		}
		return nil, "", err
	}

	f, err := cache.Get(remote)
	switch err {
	case fs.ErrorIsFile:
		return f, path.Base(fsPath), nil
	case nil:
		return f, "", nil
	default:
		cntRrr := fs.CountError(err)
		if cntRrr != nil {
			return nil, "", cntRrr
		}

		return nil, "", err
	}
}

// TODO: extract common part from the functions below

// Downloads files from connection specific storage to the local filesystem.
// Non-empty "remotePath" overrides the connection URI path.
// For example, we have the following bucket structure:
//    - /data/text.txt
// A connection has "gs://bucket-name/data" URI structure.
// The function with following parameters localPath="data/sync_dir/" remotePath="" downloads
// file to the "data/sync_dir/text.txt" location.
// The function with following parameters
// localPath="data/sync_dir/" remotePath="/data/text.txt" downloads
// file to the "data/sync_dir/text.txt" location.
// The function with following parameters
// localPath="data/sync_dir/renamed-text.txt" remotePath="/data/text.txt" downloads
// file to the "data/sync_dir/text.txt" location.
// If remotePath point to a directory then localPath must point to a directory too.
func (os *ObjectStorage) Download(localPath, remotePath string) error {
	if len(localPath) == 0 {
		return errors.New("local path is empty")
	}

	localDir, localFileName := path.Split(localPath)
	if localDir == "" {
		localDir = "./"
	}

	localFs, err := fs.NewFs(localDir)
	if err != nil {
		return err
	}

	if len(remotePath) == 0 {
		remotePath = os.RemoteConfig.Path
		log.Info("remotePath is empty, using path from remote's config", "remotePath", remotePath)
	}
	remoteWithPath := path.Join(os.RemoteConfig.FsName, remotePath)
	log.Info("Joined remote name and path", "remoteWithPath", remoteWithPath)
	remoteFs, srcFileName, err := newFsFile(remoteWithPath)
	if err != nil {
		return err
	}

	if srcFileName == "" {
		if localFileName != "" {
			return fmt.Errorf(
				"local path (%s) points to a file, but remote path (%s) points to a dir",
				localPath, remotePath,
			)
		}

		log.Info("Download directory from the remote bucket",
			"local_dir", localDir, "remote_dir", remotePath,
		)
		return sync.CopyDir(context.Background(), localFs, remoteFs, copyEmptySrcDirs)
	}

	if len(localFileName) == 0 {
		localFileName = srcFileName
	}

	log.Info("Download the file from the remote bucket",
		"local_file", localPath, "remote_file", remotePath,
	)
	return operations.CopyFile(context.Background(), localFs, remoteFs, localFileName, srcFileName)
}

// Uploads files to connection specific storage from the local filesystem.
// Non-empty "remotePath" overrides the connection URI path.
// For example, we have the following local filesystem structure:
//    - /data/text.txt
// A connection has "gs://bucket-name/sync_dir/" URI structure.
// The function with following parameters localPath="data/" remotePath="" downloads
// file to the "gs://bucket-name/sync_dir/text.txt" bucket location.
// The function with following parameters localPath="data/" remotePath="/sync_dir/" downloads
// file to the "gs://bucket-name/sync_dir/text.txt" bucket location.
// The function with following parameters
// localPath="data/text.txt" remotePath="/sync_dir/renamed-text.txt" downloads
// file to the "gs://bucket-name/sync_dir/renamed-text.txt" bucket location.
// If localPath point to a directory then remotePath must point to a directory too.
func (os *ObjectStorage) Upload(localPath, remotePath string) error {
	if len(localPath) == 0 {
		return errors.New("local path is empty")
	}

	if len(remotePath) == 0 {
		remotePath = os.RemoteConfig.Path
	}
	remoteDir, remoteFileName := path.Split(remotePath)
	rFs, err := fs.NewFs(os.RemoteConfig.FsName + remoteDir)
	if err != nil {
		return err
	}

	lFs, srcFileName, err := newFsFile(localPath)
	if err != nil {
		return err
	}

	if srcFileName == "" {
		if remoteFileName != "" {
			return fmt.Errorf("remote path (%s) points to a file, but local path (%s) points to a dir", remotePath, localPath)
		}

		log.Info("Upload the local directory to the remote bucket",
			"local_dir", localPath, "remote_dir", remotePath,
		)
		return sync.CopyDir(context.Background(), rFs, lFs, copyEmptySrcDirs)
	}

	if len(remoteFileName) == 0 {
		remoteFileName = srcFileName
	}

	log.Info("Upload the file to the remote bucket",
		"local_file", localPath, "remote_file", remotePath,
	)
	return operations.CopyFile(context.Background(), rFs, lFs, remoteFileName, srcFileName)
}
