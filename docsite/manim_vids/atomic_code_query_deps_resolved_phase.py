from manim import DOWN, Text, UP 
from manim.animation.creation import Create

from code_video import AutoScaled, CodeScene, SequenceDiagram
from code_video.widgets import DEFAULT_FONT


class SequenceDiagramScene(CodeScene):
    def construct(self):
        diagram = AutoScaled(SequenceDiagram())
        overlord, atomic_query = diagram.add_objects("Overlord", "Pasta shop Query Code")

        overlord.note("I have water, flour and salt from past errands...")
        overlord.to(atomic_query, "Would you be able to give me Pasta?")
        atomic_query.to(overlord, "yes, I can")
        overlord.note("Ok")

        title = Text("Query all deps solved", font=DEFAULT_FONT, size=0.8)
        title.to_edge(UP)
        self.add(title)
        diagram.next_to(title, DOWN)

        self.play(Create(diagram))
        for interact in diagram.get_interactions():
            self.wait(1)
            self.play(Create(interact))

        self.wait(15)

