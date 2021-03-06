function NewBufferLineReader(buf) {

	let o = Object.create(null);

	o.buf = buf;
	o.index = 0;

	o.next = function() {			// Return next line in the buffer, or null if finished.

		if (this.index >= this.buf.length) {
			return null;
		}

		let i = this.index;

		while (1) {

			if (this.buf[i] === 10 || i >= this.buf.length - 1) {
				let s;
				if (i > 0 && this.buf.length[i - 1] === 13) {
					s = this.buf.slice(this.index, i - 1).toString();		// Discard /r/n
				} else {
					s = this.buf.slice(this.index, i).toString();			// Discard /n
				}
				this.index = i + 1;
				return s;
			}

			i++;
		}
	};

	o.lines = function() {			// Remaining lines only, as an array.

		let lines = [];

		while (1) {

			let s = this.next();

			if (typeof s !== "string") {
				return lines;
			}

			lines.push(s);
		}
	};

	return o;
}
