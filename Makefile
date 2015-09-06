VC=iverilog
SIM=vvp

alu_test: alu.v alu_tb.v
	$(VC) $^ -o $@
	$(SIM) $@

clean:
	rm -rf *.vcd
	rm -rf *_test
