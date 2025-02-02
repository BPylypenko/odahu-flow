*** Variables ***
${RES_DIR}              ${CURDIR}/resources/reuters_classifier
${LOCAL_CONFIG}         odahuflow/config_e2e_mlflow_reuters_classifier
${REUTERS_CLASSIFIER}              test-e2e-reuters-classifier

*** Settings ***
Documentation       Check reuters classifier model
Test Timeout        120 minutes
Variables           ../../load_variables_from_profiles.py    ${CLUSTER_PROFILE}
Resource            ../../resources/keywords.robot
Library             Collections
Library             odahuflow.robot.libraries.utils.Utils
Library             odahuflow.robot.libraries.model.Model
Library             odahuflow.robot.libraries.examples_loader.ExamplesLoader  https://raw.githubusercontent.com/odahu/odahu-examples  ${EXAMPLES_VERSION}
Suite Setup         E2E test setup  ${REUTERS_CLASSIFIER}
Suite Teardown      E2E test teardown  ${REUTERS_CLASSIFIER}
Force Tags          e2e  reuters-classifier  cli

*** Test Cases ***
Reuters classifier model
    Download file  mlflow/tensorflow/reuters_classifier/odahuflow/training.odahuflow.yaml  ${RES_DIR}/training.odahuflow.yaml
    Download file  mlflow/tensorflow/reuters_classifier/odahuflow/packaging.odahuflow.yaml  ${RES_DIR}/packaging.odahuflow.yaml
    Download file  mlflow/tensorflow/reuters_classifier/odahuflow/deployment.odahuflow.yaml  ${RES_DIR}/deployment.odahuflow.yaml
    Download file  mlflow/tensorflow/reuters_classifier/odahuflow/request.json  ${RES_DIR}/request.json

    Run example model  ${REUTERS_CLASSIFIER}  ${RES_DIR}
