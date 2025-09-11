// src/index.ts
import dotenv from "dotenv";
import { launchSinglePage } from './browser/browser';
import { scrapeTable } from "./browser/scrape";
import { Book } from "./types/table";

// Use dotenv only if not in production
const isProduction = process.env.ENVIRONMENT === "production";
if (!isProduction) {
    dotenv.config();
}

async function main(): Promise<Book[]> {
    console.log('🌱 Starting web scraping automation...');

    const { browser, page } = await launchSinglePage(isProduction);

    const books: Book[] = await scrapeTable(page);

    console.log(`📋 Found ${books.length} books in total`);

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

// Run as ECS task
(async () => {
    try {
        const books = await main();
        console.log('✅ Web scraping automation completed');
        console.log(`Total books scraped: ${books.length}`);
        process.exit(0); // success
    } catch (err) {
        console.error('❌ Web scraping failed:', err);
        process.exit(1); // failure
    }
})();
