import asyncio

from temporalio.client import Client

from workflows import YourWorkflow, ChildWorkflowName

async def main():
    client = await Client.connect("localhost:7233")
    
    handle = await client.execute_workflow(
        YourWorkflow.run, "Child Workflow: ChildWorkflowName", id="my-workflow-id", task_queue="my-task-queue"
    )
    
    await handle.signal(YourWorkflow.Hello, "argument_Hello")
    result_Bye = await handle.query(YourWorkflow.Bye)
    print(f"Query result: {result_Bye}")


    print(f"Result: {handle}")


if __name__ == "__main__":
    asyncio.run(main())

