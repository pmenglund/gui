import { expect, test } from "@playwright/test";

test("tabs render the correct initial server state without JavaScript", async ({ browser }) => {
  const context = await browser.newContext({ javaScriptEnabled: false });
  const page = await context.newPage();

  await page.goto("http://127.0.0.1:8080/interactive");

  await expect(page.getByRole("tab", { name: "Overview" })).toHaveAttribute("aria-selected", "true");
  await expect(page.getByRole("tab", { name: "Overview" })).toHaveAttribute("tabindex", "0");
  await expect(page.getByRole("tab", { name: "Activity" })).toHaveAttribute("aria-selected", "false");
  await expect(page.getByRole("tab", { name: "Activity" })).toHaveAttribute("tabindex", "-1");
  await expect(page.locator("[data-ui-content][data-ui-target='overview']")).toBeVisible();
  await expect(page.locator("[data-ui-content][data-ui-target='activity']")).toBeHidden();

  await context.close();
});

test("interactive widgets respond to clicks and keyboard controls", async ({ page }) => {
  await page.goto("/interactive");

  await page.getByRole("button", { name: "Open dialog" }).click();
  await expect(page.getByText("The dialog content stays in place in the DOM and relies on focus restoration instead of portals.")).toBeVisible();
  await page.keyboard.press("Escape");
  await expect(page.getByText("The dialog content stays in place in the DOM and relies on focus restoration instead of portals.")).toBeHidden();

  await page.getByRole("button", { name: "Quick actions" }).click();
  await expect(page.getByText("View logs")).toBeVisible();
  await page.mouse.click(5, 5);
  await expect(page.getByText("View logs")).toBeHidden();

  await page.getByRole("tab", { name: "Overview" }).focus();
  await page.keyboard.press("ArrowRight");
  await expect(page.getByRole("tab", { name: "Activity" })).toBeFocused();
  await expect(page.getByText("The tabs runtime handles keyboard navigation with arrow keys.")).toBeVisible();

  await page.getByRole("button", { name: "Show toast" }).click();
  await expect(page.getByText("Settings saved")).toBeVisible();

  await page.getByRole("button", { name: "Open sheet" }).click();
  await expect(page.getByText("This panel uses the right edge of the viewport and the same close affordances.")).toBeVisible();
  await page.keyboard.press("Escape");
  await expect(page.getByText("This panel uses the right edge of the viewport and the same close affordances.")).toBeHidden();
});

test("htmx flows update fragments and swapped overlays stay interactive", async ({ page }) => {
  await page.goto("/htmx");

  await page.getByRole("button", { name: "Refresh counter fragment" }).click();
  await expect(page.getByText("Counter update 2")).toBeVisible();

  const email = page.getByLabel("Email");
  await email.click();
  await email.type("ops@example.com");
  await expect(page.locator("#email-hint")).toHaveText("Looks good for a company email.");

  await page.getByRole("link", { name: "2" }).last().click();
  await expect(page.getByText("Search")).toBeVisible();

  await page.getByRole("button", { name: "Load interactive fragment" }).click();
  await page.getByRole("button", { name: "Open swapped dialog" }).click();
  await expect(page.getByText("Delegated runtime hooks mean the client code does not need a manual rebind.")).toBeVisible();
});
