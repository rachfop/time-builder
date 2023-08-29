from temporalio import workflow
from datetime import timedelta

with workflow.unsafe.imports_passed_through():
    
    from activities import YourParams
    

@workflow.defn
class YourWorkflow:
    @workflow.run
    async def run(self, name: str) -> str:
        
        return await workflow.execute_child_workflow(
            ChildWorkflowName.run,
            YourParams("Hello", name),
            id="child-workflow-id",
            schedule_to_close_timeout=timedelta(seconds=5),
        )
        
    @workflow.signal
    async def Hello(self):
        pass
    @workflow.query
    async def Bye(self):
        pass



@workflow.defn
class ChildWorkflowName:
    @workflow.run
    async def run(self, name: str) -> str:
        return await workflow.execute_activity(
            say_hello,
            YourParams("Hello", name),
            schedule_to_close_timeout=timedelta(seconds=5),
        )
