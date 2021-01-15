*** Variables ***
${LOCAL_CONFIG}         odahuflow/config_common_verbose
${TEST_VALUE}           test

*** Settings ***
Documentation       Verbose cli config command
Resource            ../../resources/keywords.robot
Force Tags          cli  common
Suite Setup         Set Environment Variable  ODAHUFLOW_CONFIG  ${LOCAL_CONFIG}
Suite Teardown      Remove file  ${LOCAL_CONFIG}

*** Keywords ***
Should contain stacktrace
    [Arguments]  ${message}

    Should Contain  ${message}  Traceback (most recent call last)

Should contain debug logs
    [Arguments]  ${message}

    Should Contain  ${message}  - DEBUG -
    Should Contain  ${message}  - ERROR - Exception occurs during CLI invocation

Should contain error message
    [Arguments]  ${message}

    Should contain  ${message}  Error:

Help Option
    [Arguments]  ${command}
    ${cmd_h}  StrictShell  ${command} -h
    Should contain  ${cmd_h.stdout}  Usage: odahuflowctl
    Should not contain  ${cmd_h.stderr}  Error

    ${cmd_help}  StrictShell  ${command} --help
    Should contain  ${cmd_help.stdout}  Usage: odahuflowctl
    Should not contain  ${cmd_help.stderr}  Error

*** Test Cases ***
The verbose flag is disabled
    [Tags]  verbose
    ${res}=  Shell  odahuflowctl --non-verbose conn get
    Should not be equal  ${res.rc}  ${0}
    Should contain error message  ${res.stdout}
    Should contain  ${res.stdout}  For more information rerun command with --verbose flag

    Run Keyword And Expect Error  *  Should contain stacktrace  ${res.stderr}
    Run Keyword And Expect Error  *  Should contain debug logs  ${res.stderr}
    Run Keyword And Expect Error  *  Should contain stacktrace  ${res.stdout}
    Run Keyword And Expect Error  *  Should contain debug logs  ${res.stdout}

The verbose flag is disabled by default
    [Tags]  verbose
    ${res}=  Shell  odahuflowctl conn get
    Should not be equal  ${res.rc}  ${0}
    Should contain error message  ${res.stdout}
    Should contain  ${res.stdout}  For more information rerun command with --verbose flag

    Run Keyword And Expect Error  *  Should contain stacktrace  ${res.stderr}
    Run Keyword And Expect Error  *  Should contain debug logs  ${res.stderr}
    Run Keyword And Expect Error  *  Should contain stacktrace  ${res.stdout}
    Run Keyword And Expect Error  *  Should contain debug logs  ${res.stdout}

The verbose flag is enabled
    [Tags]  verbose
    ${res}=  Shell  odahuflowctl --verbose conn get
    Should not be equal  ${res.rc}  ${0}
    Should contain error message  ${res.stdout}

    Should contain stacktrace  ${res.stderr}
    Should contain debug logs  ${res.stderr}
    Run Keyword And Expect Error  *  Should contain stacktrace  ${res.stdout}
    Run Keyword And Expect Error  *  Should contain debug logs  ${res.stdout}

The help option (-h, --help)
    [Tags]  help-option
    [Template]  Help Option
    odahuflowctl --verbose conn get
    odahuflowctl --verbose train
    odahuflowctl packaging
    odahuflowctl local training
    odahuflowctl --verbose local train run
    odahuflowctl local pack
    odahuflowctl --verbose pi --url http edit --id not-exist