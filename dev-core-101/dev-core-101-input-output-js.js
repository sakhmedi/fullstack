const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
});

readline.question("Enter your name: ", userName => {
    readline.question("Enter your age: ", ageInput => {
        let age = parseInt(ageInput);
        let currentYear = new Date().getFullYear();
        let yearHundred = currentYear + (100 - age);
        
        console.log(`Hello, ${userName}! You will turn 100 years old in the year ${yearHundred}.`);
        readline.close();
    });
});
