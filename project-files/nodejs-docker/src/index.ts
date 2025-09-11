// index.js
console.log("ECS task started");

const now = new Date().toISOString();
console.log(`Task execution started at ${now}`);

setTimeout(() => {
    const doneAt = new Date().toISOString();
    console.log(`Task finished at ${doneAt}`);
    process.exit(0); // exit successfully
}, 5000);
