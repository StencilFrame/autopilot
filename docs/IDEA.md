# Embracing Gradual Automation: My Journey to Building DoNot

Hey everyone!

I want to share something that’s been brewing in my mind for a while—the concept of gradual automation and how it inspired me to create DoNot, a tool aimed at transforming the way we approach automating our workflows.

## My Love Affair with Automation

For as long as I can remember, I’ve been fascinated by automation. Throughout my career, I’ve poured countless hours into automating everything I could get my hands on. Whether it was streamlining DevOps processes, enhancing CI/CD pipelines, tackling mundane repetitive tasks, or improving the overall developer experience, I was always on the lookout for ways to make things run smoother and more efficiently.

But here’s the thing—while automation was my passion, I often found myself grappling with the limitations of existing tools. Most of them were designed to handle tasks in an all-or-nothing fashion: either you automate the entire process from start to finish or you stick with manual execution. There was little room for a middle ground, a space where you could gradually introduce automation without overhauling your entire workflow overnight.

## The Eureka Moment: Discovering Gradual Automation

Back in 2019, I stumbled upon an article by Dan Slimmon titled “[Do Nothing Scripting: The Key to Gradual Automation](https://blog.danslimmon.com/2019/07/15/do-nothing-scripting-the-key-to-gradual-automation)”. Let me tell you, it was a game-changer. The idea struck a chord deep within me—it was brilliant, innovative, insightful, and downright genius.

Dan Slimmon emphasized the importance of gradual automation—a strategy where you start by defining all steps as manual workflows and then decide which parts make the most sense to automate. You automate step by step, refining and enhancing the process until you either fully automate it or leave just a few manual touches. This approach resonated with me because it addressed the very frustration I felt with existing tools.

## Why Gradual Automation Makes Sense

1. Incremental Automation

    Automating an entire workflow in one go can be overwhelming. There’s so much to consider—dependencies, error handling, and ensuring that every step works seamlessly together. With incremental automation, you can focus on automating one step at a time, ensuring each part works perfectly before moving on to the next.

2. Maintains Control

    Gradual automation gives you control over the process. You can decide which parts to automate based on priority and impact, rather than being forced into a complete automation overhaul. This way, you retain flexibility and can adapt as your workflows evolve.

3. Builds Confidence

    Starting small allows you to build confidence in your automation efforts. As you successfully automate individual steps, you’ll gain the assurance needed to tackle more complex parts of your workflow.

4. Optimizes Resources

    Not all tasks are created equal. Some are ripe for automation, while others benefit from human intuition and oversight. Gradual automation helps you identify and prioritize the tasks that truly need automation, ensuring that you invest your resources wisely.

5. Reduces Errors and Cognitive Load

    Manual runbooks are a breeding ground for human error and inconsistency. I can't remember how many times I've skipped a step executed a task incorrectly, or lost track of the progress due to distractions, leading to wasted time and effort. Having a tool that wraps manual steps and tracks the progress of execution ensures that you don't miss a step. Also, allows you to resume from the last step in case of an interruption.

6. Encourages Collaboration

    Having ability to transfer a runbook in-progress to another team member is a game-changer. It allows you to share the workload and collaborate more effectively, ensuring that everyone is on the same page and can contribute to the automation process.

## Introducing DoNot: Your Partner in Gradual Automation

[Do Not]hing Scripting Automation.

Inspired by Dan Slimmon’s insights, I embarked on developing DoNot—a tool designed to facilitate gradual automation. The core idea behind DoNot is to allow users to define runbooks in Markdown or YAML formats, supporting both manual and shell steps. This flexibility means you can start with a fully manual workflow and incrementally automate the parts that make the most sense for your specific needs.

### What DoNot Brings to the Table
* Supports Manual and Shell Steps: Begin by outlining your entire workflow with manual steps. As you identify automation opportunities, seamlessly convert those steps into shell commands, reducing manual intervention over time.
* Manual steps are first-class citizens: DoNot treats manual steps with the same importance as shell steps, ensuring that you have complete control over the automation process.
* Simple to Use: Define your runbook in Markdown or YAML, and let DoNot handle the rest. The tool is designed to be intuitive and user-friendly, making it easy to get started with gradual automation.
* Extensible and Customizable: DoNot is designed to grow with your needs. As you become more comfortable with automation, you can expand your runbooks to include more complex step types, conditional logic, and other advanced features. (Work in progress)
* Open Source: DoNot is an open-source project, meaning you can contribute to its development, suggest new features, or customize it to suit your specific requirements. And esspecially you can fix bugs and issues you encounter. :-)

### How It Works
1.	Define Your Runbook: Start by listing out all the steps in your workflow using Markdown or YAML. Initially, all steps can be manual, allowing you to map out the entire process clearly.
2.	Identify Automation Targets: Review your runbook and pinpoint the steps that are repetitive or time-consuming—the ideal candidates for automation.
3.	Automate Incrementally: Convert selected manual steps into shell commands. Test each automated step thoroughly before moving on to the next, ensuring reliability and stability.
4.	Iterate and Enhance: Continue the process of automating steps one by one, gradually transforming your workflow into a more efficient, automated system.

## A Real-World Example

Let me share a quick example of how gradual automation with DoNot can transform a workflow:

Scenario: you have to build a new feature for your application: self serving onboarding of the new workspace for user. The process involves creating records in the SQL database and setting up the index in the NoSQL database.

You are not sure if the feature will be successful, so you don't want to invest too much time in automation upfront (build what matters), or maybe you have another high-priority task that requires your attention. But you also want to ensure that the process is documented, repeatable and error-free. So you decide that you can onboard workspaces manually using documentation until you see the feature gaining traction.

Here's how you could define the runbook for this process:
~~~
# Onboard New Workspace

Requirements:
    - user_email:
    - workspace_name:
    - MYSQL connection string:
    - NOSQL connection string:

Variables:
    - workspace_id:
    - user_id:

1. Connect to the SQL database to create a new workspace record
    [Provide instructions for how to connect to the right database]
2. Create a new workspace record in the SQL database
    ```
    INSERT INTO WORKSPACES (name) values ('{workspace_name}');
    ```
3. Get the newly created workspace ID from the SQL database as {workspace_id}
    ```
    SELECT id FROM WORKSPACES WHERE name = '{workspace_name}';
    ```
4. Get user ID from the SQL database as {user_id}
    ```
    SELECT id FROM USERS WHERE email = '{user_email}';
    ```
5. Add user to the workspace
    ```
    INSERT INTO WORKSPACE_USERS (workspace_id, user_id) values ({workspace_id}, {user_id});
    ```
6. Connect to the NoSQL database to create an index for the new workspace
    [Provide instructions for how to connect to the right database]
7. Create an index in the NoSQL database
    ```
    CREATE INDEX `event_index_<workspace_id>` IF NOT EXISTS
    ...
    ```
~~~

NOTE: having this runbook as a document is already a huge improvement over having nothing, but executing it is still error-prone as you have to keep track of the progress and all variables. Step 3 and 7 share the same variable `{workspace_id}`. If you forget to update it in step 7, you will create an index for the wrong workspace.

With DoNot, you can start automating the steps that are most time-consuming or error-prone. For example, you could automate steps 2, 3, 4, and 5, while leaving steps 1, 6, and 7 as manual tasks. This way, you can ensure that the critical parts of the process are automated, while still maintaining control over the entire workflow.

DoNot will track the progress of the execution and all variables, so you don't have to worry about missing a step or updating the wrong variable. It can also show you all the metadata about the execution, so you can easily resume from the last step in case of an interruption.

As the feature gains traction and you see the need for more automation, you can gradually convert the remaining manual steps into shell commands, making the process more efficient and scalable. And eventually, you can move automation to the fully functional feature of your application.

## Why I Believe in Gradual Automation

Over the years, I’ve seen teams struggle with the rigidity of full automation and the inefficiency of entirely manual processes. Gradual automation offers a balanced approach, marrying the best of both worlds. It ensures that automation enhances productivity without introducing unnecessary complexity or loss of control.

By leveraging DoNot, you can tailor your automation journey to fit your unique needs, ensuring that each step is optimized for maximum efficiency and reliability. It’s about making smart, incremental changes that add up to significant improvements over time.

## Join the Journey

I’m incredibly excited about the potential of DoNot to redefine automation practices. But I can’t do it alone—I need your feedback! Whether you’re a seasoned DevOps engineer, an IT professional, or someone curious about workflow automation, your insights are invaluable.

Check out DoNot on GitHub: [https://github.com/StencilFrame/donot](https://github.com/StencilFrame/donot)

I’d love to hear your thoughts:
* Do you find the concept of gradual automation valuable?
* What do you like or dislike about the current features?
* Are there any features you’re missing that would make DoNot more useful for your daily tasks?
* What would encourage you to use it regularly in your workflow?

Your feedback will help shape the future of DoNot, ensuring it meets the real-world needs of users like you. Together, let’s make automation smarter, more manageable, and truly beneficial.

Thanks for taking the time to read this, and here’s to smarter automation!

Cheers,  
Mindaugas