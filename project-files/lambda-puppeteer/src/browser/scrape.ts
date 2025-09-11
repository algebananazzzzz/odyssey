import { Page } from "puppeteer-core";
import { logStep } from "../helpers";
import { Book } from "../types/table";

export async function scrapeTable(page: Page): Promise<Book[]> {
    let step = performance.now();
    const url = "https://books.toscrape.com/";

    await page.goto(url, { waitUntil: "networkidle2" });
    logStep(`🌱 Navigate to search page: ${url}`, step);

    // Race between jobs or empty state
    step = performance.now();
    const books: Book[] = await page.$$eval("article.product_pod", (items) =>
        items.map((el) => {
            const title = el.querySelector("h3 a")?.getAttribute("title") || "";
            const url = el.querySelector("h3 a")?.getAttribute("href") || "";
            const price = el.querySelector(".price_color")?.textContent.trim() || "";
            const availability =
                el.querySelector(".instock.availability")?.textContent.trim() || "";
            const rating =
                el.querySelector("p.star-rating")?.classList[1] || "None";

            return {
                title,
                url,
                price,
                availability,
                rating,
            };
        })
    );

    logStep("Finished scraping books", step);

    return books;
}