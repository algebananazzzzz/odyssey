import dotenv from "dotenv";


const isProduction = process.env.ENVIRONMENT === "production";
if (!isProduction) {
    dotenv.config();
}

export const handler = async (event) => {
    console.log("Event received from EventBridge:", JSON.stringify(event, null, 2));

    const now = new Date().toISOString();
    console.log(`Lambda executed at ${now}`);

    return {
        statusCode: 200,
        body: JSON.stringify({ message: "Event processed", time: now })
    };
};