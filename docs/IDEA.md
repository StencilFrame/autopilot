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

1. Reduces Overwhelm

    Automating an entire workflow in one go can be overwhelming. There’s so much to consider—dependencies, error handling, and ensuring that every step works seamlessly together. By taking it slow, you can focus on automating one step at a time, ensuring each part works perfectly before moving on to the next.

2. Maintains Control

    Gradual automation gives you control over the process. You can decide which parts to automate based on priority and impact, rather than being forced into a complete automation overhaul. This way, you retain flexibility and can adapt as your workflows evolve.

3. Builds Confidence

    Starting small allows you to build confidence in your automation efforts. As you successfully automate individual steps, you’ll gain the assurance needed to tackle more complex parts of your workflow.

4. Optimizes Resources

    Not all tasks are created equal. Some are ripe for automation, while others benefit from human intuition and oversight. Gradual automation helps you identify and prioritize the tasks that truly need automation, ensuring that you invest your resources wisely.

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

Scenario: You’re managing the setup of new development environments—a task that involves several repetitive steps.
	1.	Step 1: Verify system prerequisites.
	2.	Step 2: Run a script to install necessary software.
	3.	Step 3: Configure environment variables.
	4.	Step 4: Execute post-configuration scripts.

With DoNot, you can start by automating Step 2. Once you’re confident that the installation script runs flawlessly, you can move on to automating Step 4. This way, you gradually reduce manual workload without disrupting the entire setup process.

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