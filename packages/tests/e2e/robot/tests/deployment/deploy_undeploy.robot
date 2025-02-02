*** Variables ***
${RES_DIR}              ${CURDIR}/resources
${LOCAL_CONFIG}        odahuflow/config_deployment_dep_undep
${MD_SIMPLE_MODEL}     simple-model-dep-undep

*** Settings ***
Documentation       Test model deployment through CLI for normal operational processes
Test Timeout        20 minutes
Resource            ../../resources/keywords.robot
Resource            ../../resources/variables.robot
Variables           ../../load_variables_from_profiles.py    ${CLUSTER_PROFILE}
Library             odahuflow.robot.libraries.k8s.K8s  ${ODAHUFLOW_DEPLOYMENT_NAMESPACE}
Library             odahuflow.robot.libraries.utils.Utils
Library             Collections
Suite Setup         Run Keywords  Set Environment Variable  ODAHUFLOW_CONFIG  ${LOCAL_CONFIG}  AND
...                               Login to the api and edge  AND
...                               Cleanup resources
Suite Teardown      Run keywords  Cleanup resources  AND
...                 Remove File  ${LOCAL_CONFIG}

Force Tags        cli  deployment

*** Keywords ***
Cleanup resources
    StrictShell  odahuflowctl --verbose dep delete --id ${MD_SIMPLE_MODEL} --ignore-not-found

File not found
    [Arguments]  ${command}
        ${res}=  Shell  odahuflowctl --verbose dep ${command} -f wrong-file
                 Should not be equal  ${res.rc}  ${0}
                 Should contain       ${res.stderr}  Resource file 'wrong-file' not found

*** Test Cases ***
Check deploy procedure
    [Documentation]  Try to deploy dummy model through API console
    [Teardown]  Cleanup resources
    Run API deploy from model packaging  ${MP_SIMPLE_MODEL}  ${MD_SIMPLE_MODEL}  ${RES_DIR}/simple-model.deployment.odahuflow.yaml

    Check model started  ${MD_SIMPLE_MODEL}

Update model deployment
    [Documentation]  Check model deployment update
    [Teardown]  Cleanup resources
    Run API deploy from model packaging  ${MP_SIMPLE_MODEL}  ${MD_SIMPLE_MODEL}  ${RES_DIR}/simple-model.deployment.odahuflow.yaml
    Check model started  ${MD_SIMPLE_MODEL}

    ${res}=  StrictShell  odahuflowctl --verbose model invoke --md ${MD_SIMPLE_MODEL} --json '{"columns": ["a","b"],"data": [[1.0,2.0]]}'
    ${actual_response}=    Evaluate     json.loads("""${res.stdout}""")    json
    ${expected_response}=   evaluate  {'prediction': [[42]], 'columns': ['result']}
    Dictionaries Should Be Equal    ${actual_response}    ${expected_response}

    Run API apply from model packaging  ${MP_COUNTER_MODEL}  ${MD_SIMPLE_MODEL}  ${RES_DIR}/updated-simple-model.deployment.odahuflow.yaml
    Check model started  ${MD_SIMPLE_MODEL}

    # Check new REST API
    ${res}=  StrictShell  odahuflowctl --verbose model invoke --md ${MD_SIMPLE_MODEL} --json '{"columns": ["a","b"],"data": [[1.0,2.0]]}'
    ${actual_response}=    Evaluate     json.loads("""${res.stdout}""")    json
    ${expected_response}=   evaluate  {'prediction': [[1]], 'columns': ['result']}
    Dictionaries Should Be Equal    ${actual_response}    ${expected_response}

    # Check new resources
    ${res}=  StrictShell  odahuflowctl dep get --id ${MD_SIMPLE_MODEL} -o jsonpath='[*].status.deployment'
    ${model_deployment}=  Get model deployment  ${res.stdout}  ${ODAHUFLOW_DEPLOYMENT_NAMESPACE}
    LOG  ${model_deployment}

    ${model_resources}=  Set variable  ${model_deployment.spec.template.spec.containers[0].resources}
    Should be equal  333m  ${model_resources.limits["cpu"]}
    Should be equal  333Mi  ${model_resources.limits["memory"]}
    Should be equal  222m  ${model_resources.requests["cpu"]}
    Should be equal  222Mi  ${model_resources.requests["memory"]}

    # Check new number of replicas
    ${res}=  StrictShell  odahuflowctl dep get --id ${MD_SIMPLE_MODEL} -o jsonpath='[*].status.replicas'
    should be equal as integers  ${2}  ${res.stdout}

Deploy with custom memory and cpu
    [Documentation]  Deploy with custom memory and cpu
    [Teardown]  Cleanup resources
    Run API deploy from model packaging  ${MP_SIMPLE_MODEL}  ${MD_SIMPLE_MODEL}  ${RES_DIR}/custom-resources.deployment.odahuflow.yaml

    ${res}=  StrictShell  odahuflowctl dep get --id ${MD_SIMPLE_MODEL} -o jsonpath='[*].status.deployment'
    ${model_deployment}=  Get model deployment  ${res.stdout}  ${ODAHUFLOW_DEPLOYMENT_NAMESPACE}
    LOG  ${model_deployment}

    ${model_resources}=  Set variable  ${model_deployment.spec.template.spec.containers[0].resources}
    Should be equal  333m  ${model_resources.limits["cpu"]}
    Should be equal  333Mi  ${model_resources.limits["memory"]}
    Should be equal  222m  ${model_resources.requests["cpu"]}
    Should be equal  222Mi  ${model_resources.requests["memory"]}

Check setting of default resource values
    [Documentation]  Deploy setting of default resource values
    [Teardown]  Cleanup resources
    Run API deploy from model packaging  ${MP_SIMPLE_MODEL}  ${MD_SIMPLE_MODEL}  ${RES_DIR}/simple-model.deployment.odahuflow.yaml

    ${res}=  StrictShell  odahuflowctl dep get --id ${MD_SIMPLE_MODEL} -o jsonpath='[*].status.deployment'
    ${model_deployment}=  Get model deployment  ${res.stdout}  ${ODAHUFLOW_DEPLOYMENT_NAMESPACE}
    LOG  ${model_deployment}

    ${model_resources}=  Set variable  ${model_deployment.spec.template.spec.containers[0].resources}
    Should be equal  250m  ${model_resources.limits["cpu"]}
    Should be equal  256Mi  ${model_resources.limits["memory"]}
    Should be equal  125m  ${model_resources.requests["cpu"]}
    Should be equal  128Mi  ${model_resources.requests["memory"]}

File with entitiy not found
    [Documentation]  Invoke Model Deployment commands with not existed file
    [Template]  File not found
    command=create
    command=edit
