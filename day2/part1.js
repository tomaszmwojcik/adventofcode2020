// Input format
// n-m a: pass
// n - lower bound for a occurrences of [a] in [pass]
// m - upper bound -||-
// Output: how many passwords are valid according to lower/upper bound rules

const { match } = require('assert');
const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
  terminal: false
});

const inputPattern = /([0-9]+)-([0-9]+) (.): (.+)/;
let matchingPasswords = 0;

rl.on('line', (input) => {
    let parsed = input.match(inputPattern);
    const lower = parseInt(parsed[1]);
    const upper = parseInt(parsed[2]);
    const letter = parsed[3];
    const password = parsed[4];
    let counter = 0;
    for(const ch of password) {
        if(ch === letter) {
            counter++;
        }
    }
    if(counter >= lower && counter <= upper) {
        matchingPasswords++;
    }
  });

rl.on('close', () => {
    console.log(`There are ${matchingPasswords} matching passwords.`);
});