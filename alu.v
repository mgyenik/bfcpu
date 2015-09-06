module alu(op, a, b, out);
	/* Width of the inputs/outputs. */
	parameter DATA_WIDTH = 8;

	input wire op;
	input wire[DATA_WIDTH-1:0] a;
	input wire[DATA_WIDTH-1:0] b;
	output wire[DATA_WIDTH-1:0] out;

	wire[DATA_WIDTH-1:0] add;
	wire[DATA_WIDTH-1:0] sub;

	assign add = a + b;
	assign sub = a - b;

	/* 0 is add, 1 is subtract. */
	assign out = op ? sub : add;

endmodule
