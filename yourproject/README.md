# YourWorkflow project

<!-- Project description. -->

To run, first start the Temporal server:

```command
temporal server start-dev
```

Then, in another terminal, run the worker and workflow:

```command
poetry run python hello/run_worker.py
poetry run python hello/run_workflow.py
```
