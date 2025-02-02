#!/usr/bin/env bash
set -e

AIRFLOW_NAMESPACE=airflow
AIRFLOW_WEB_CONTAINER_NAME=airflow-web
KUBECTL="kubectl --request-timeout 10s"

function ReadArguments() {
  if [[ $# == 0 ]]; then
    echo "ERROR: Options not specified! Use -h for help!"
    exit 1
  fi

  while [[ $# -gt 0 ]]; do
    case "$1" in
    -h | --help)
      echo "$0 - Launch and wait finish of the dags"
      echo -e "Usage: $0 [OPTIONS]\n\noptions:"
      echo -e "--dags\t\tDags for testing, for example: --dags 'airflow-wine,airflow-tensorflow'"
      echo -e "-v  --verbose\t\tverbose mode for debug purposes"
      echo -e "-h  --help\t\tshow brief help"
      exit 0
      ;;
    --dags)
      export TEST_DAG_IDS_RAW="$2"
      shift 2
      ;;
    -v | --verbose)
      export VERBOSE=true
      shift
      ;;
    *)
      echo "ERROR: Unknown option: $1. Use -h for help."
      exit 1
      ;;
    esac
  done

  # Check mandatory parameters
  if [[ ! $TEST_DAG_IDS_RAW ]]; then
    echo "ERROR: dags argument must be specified. Use -h for help!"
    exit 1
  else
    IFS=',' read -r -a TEST_DAG_IDS <<< "${TEST_DAG_IDS_RAW}"
    export TEST_DAG_IDS
  fi

  if [[ $VERBOSE == true ]]; then
    set -x
  fi
}

function wait_dags_finish() {
  for i in "${!TEST_DAG_RUN_IDS[@]}"; do
    dag_run_id="${TEST_DAG_RUN_IDS[${i}]}"
    dag_id="${TEST_DAG_IDS[${i}]}"

    echo "Wait for the finishing of ${dag_id} and ${dag_run_id} its run"

    while true; do
      # Extract a dag state from the following output json.
      #[
      # ...,
      #  {
      #    "dag_id": "health_check",
      #    "dag_run_url": "/airflow/admin/airflow/graph?dag_id=health_check&execution_date=2020-09-16+09%3A55%3A00%2B00%3A00",
      #    "execution_date": "2020-09-16T09:55:00+00:00",
      #    "id": 43,
      #    "run_id": "scheduled__2020-09-16T09:55:00+00:00",
      #    "start_date": "2020-09-16T10:00:01.489937+00:00",
      #    "state": "success"
      #  }
      #]
      state=$(${KUBECTL} exec "$POD" -n "${AIRFLOW_NAMESPACE}" -c "${AIRFLOW_WEB_CONTAINER_NAME}" -- \
      curl -X GET  http://localhost:8080/airflow/api/experimental/dags/${dag_id}/dag_runs \
           -H 'Cache-Control: no-cache'   -H 'Content-Type: application/json' \
           -d "{\"run_id\": \"${dag_id}\"}" --silent \
           | jq -r -c ".[] | select(.run_id == \"${dag_run_id}\") | .state")
      echo "${state}"
      case "${state}" in
      "success")
        echo "DAG run ${dag_run_id} finished"
        break
        ;;
      "running")
        echo "DAG run ${dag_run_id} is running. Sleeping 30 sec..."
        sleep 30
        ;;
      "failed")
        echo "DAG run ${dag_run_id} failed"
        exit 1
        ;;
      *)
        echo "${state} is unknown state of the ${dag_run_id} DAG"
        exit 1
        ;;
      esac
    done
  done
}

export TEST_DAG_RUN_IDS=()
ReadArguments "$@"
POD=$(${KUBECTL} get pods -l app=airflow -l component=web -n "${AIRFLOW_NAMESPACE}" -o jsonpath='{.items[0].metadata.name}')
export POD

# Run all test dags
for dag_id in "${TEST_DAG_IDS[@]}"; do
  dag_run_id="${dag_id}-ci-$(date +%s)"
  TEST_DAG_RUN_IDS+=("${dag_run_id}")

  echo "Run the ${dag_run_id} of ${dag_id} dag"
  ${KUBECTL} exec "$POD" -n "${AIRFLOW_NAMESPACE}" -c "${AIRFLOW_WEB_CONTAINER_NAME}" -- \
  curl -X POST   http://localhost:8080/airflow/api/experimental/dags/${dag_id}/dag_runs  \
       -H 'Cache-Control: no-cache'   -H 'Content-Type: application/json'  \
       -d "{\"run_id\": \"${dag_run_id}\"}" --silent

done

wait_dags_finish
