import asyncio



from temporalio.client import Client
from temporalio.worker import Worker
from workflows import YourWorkflow, ChildWorkflowName
from activities import say_hello



async def main():
    
    client = await Client.connect("localhost:7233")
    

    
    worker = Worker(
        client, task_queue="my-task-queue", workflows=[YourWorkflow, ChildWorkflowName], activities=[say_hello]
    )
    
    await worker.run()


if __name__ == "__main__":
    asyncio.run(main())
