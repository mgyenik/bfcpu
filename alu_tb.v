module alu_tb;
	reg[15:0] a;
	reg[15:0] b;
	wire[15:0] out;
	reg op;

	/* Test a 16-bit ALU. */
	alu #(16) DUT(op, a, b, out);

	initial begin
		$dumpfile("alu_test.vcd");
		$dumpvars(0, DUT);

		/* First insert data, add output will be 0x69, subtract output will be 0x9. */
		a <= 'h39;
		b <= 'h30;

		/* First test add for 5 cycles. */
		op <= 0;

		/* Then test subtract. */
		#5 op <= 1;
		#5 $finish();
	end
endmodule
