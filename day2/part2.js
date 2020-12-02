// Input format
// n-m a: pass
// n, m - exactly one of these position in [pass] should contain [a]
// Output: how many passwords are valid 

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
    const firstPos = parseInt(parsed[1]);
    const secondPos = parseInt(parsed[2]);
    const letter = parsed[3];
    const password = parsed[4];
    let counter = 0;
    if(password[firstPos-1] === letter) {
        counter++;
    }
    if(password[secondPos-1] === letter) {
        counter++;
    }
    if(counter === 1) {
        matchingPasswords++;
    }
  });

rl.on('close', () => {
    console.log(`There are ${matchingPasswords} matching passwords.`);
});