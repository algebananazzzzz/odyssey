import dotenv from "dotenv";
import { launchSinglePage } from './browser/browser';
import { scrapeTable } from "./browser/scrape";
import { Book } from "./types/table";

const isProduction = process.env.ENVIRONMENT === "production";
if (!isProduction) {
    dotenv.config();
}

async function main(): Promise<Book[]> {
    console.log('🌱 Starting web scraping automation...');

    const { browser, page } = await launchSinglePage(isProduction);

    const books: Book[] = await scrapeTable(page);

    console.log(`📋 Found ${books.length} in total`)

    // Pretty-print all books
    books.forEach((book, i) => {
        console.log(`📖 Book #${i + 1}`);
        console.log(`   Title:        ${book.title}`);
        console.log(`   Price:        ${book.price}`);
        console.log(`   Availability: ${book.availability}`);
        console.log(`   Rating:       ${book.rating}`);
        console.log(`   URL:          ${book.url}`);
    });

    await browser.close();
    return books;
}

export const handler = async (event) => {
    const books: Book[] = await main();
    const response = {
        statusCode: 200,
        body: JSON.stringify({
            message: '✅ Web scraping automation completed',
            count: books.length,
            data: books,
        }),
    };
    return response;
};

if (!isProduction) {
    main()
}