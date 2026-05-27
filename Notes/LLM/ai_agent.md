1. what ai agent is:
    - user input via unstructured text
    - agent convert to structured instruction
    - agent do thing to achieve

2. Single responsibility principle
    - keep each agents focus on one thing
    - agent patternL:
        1. selector: summarize conversation, extract intent, decide what to do next
        2. database agent: structured filters
        3. search agent
        4. refine agent
        5. general agent: analysis, summarize
        6. update agent:
        7. launch agent: finalize step

3. do not overfeed model
    - need to pass context to agents, but need to control
    - summarize last 5, keep prompt short/precise, pass only what next agent needs

4. orchestration
    - autonomy way: each agents call another agent at will
        - cannot understand why each are being called
        - cannot know why token burnt
    - orchestration way:
        - selector to choose next agent
        - 1 specialized agent execute
        - aggregator/formatter finalizes the output
        - keep within 2-3 hop (selector -> search -> formatter, 2 hops)
            - better success rate
5. need to see what agent did and why to debug and improve
    -  LangSmith and Helicone to trace calls

6. do test

7. keep agents focus on structured criteria
    - let agent extract filters from text, let ORM to query