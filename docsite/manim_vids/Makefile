.PHONY: vids

vids: out
	./codevidgen.sh atomic_code_execution_phase.py
	./codevidgen.sh atomic_code_query_phase.py
	./codevidgen.sh atomic_code_query_deps_resolved_phase.py
	./codevidgen.sh atomic_code_query_deps_partial_phase.py
	cp media/videos/atomic_code_execution_phase/1080p60/SequenceDiagramScene.mp4 out/execution.mp4
	cp media/videos/atomic_code_query_phase/1080p60/SequenceDiagramScene.mp4 out/query.mp4
	cp media/videos/atomic_code_query_deps_partial_phase/1080p60/SequenceDiagramScene.mp4 out/partial.mp4
	cp media/videos/atomic_code_query_deps_resolved_phase/1080p60/SequenceDiagramScene.mp4 out/resolved.mp4

out:
	mkdir out

