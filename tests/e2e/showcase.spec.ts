import { expect, test } from "@playwright/test";

test("interactive widgets respond to clicks and keyboard controls", async ({ page }) => {
  await page.goto("/interactive");

  const dialogTrigger = page.getByRole("button", { name: "Open dialog" });
  const dialogSurface = page.getByRole("dialog", { name: "Review rollout" });
  await dialogTrigger.click();
  await expect(dialogSurface).toBeFocused();
  await expect(page.getByText("The dialog content stays in place in the DOM and relies on focus restoration instead of portals.")).toBeVisible();
  await page.keyboard.press("Escape");
  await expect(dialogTrigger).toBeFocused();
  await expect(page.getByText("The dialog content stays in place in the DOM and relies on focus restoration instead of portals.")).toBeHidden();
  await dialogTrigger.click();
  await expect(dialogSurface).toBeFocused();
  await dialogSurface.getByRole("button", { name: "Close" }).click();
  await expect(dialogTrigger).toBeFocused();
  await expect(page.getByText("The dialog content stays in place in the DOM and relies on focus restoration instead of portals.")).toBeHidden();

  const quickActionsTrigger = page.getByRole("button", { name: "Quick actions" });
  const viewLogs = page.getByRole("link", { name: "View logs" });
  await quickActionsTrigger.click();
  await expect(viewLogs).toBeFocused();
  await expect(viewLogs).toBeVisible();
  await page.mouse.click(5, 5);
  await expect(quickActionsTrigger).toBeFocused();
  await expect(viewLogs).toBeHidden();

  await page.getByRole("tab", { name: "Overview" }).focus();
  await page.keyboard.press("ArrowRight");
  await expect(page.getByRole("tab", { name: "Activity" })).toBeFocused();
  await expect(page.getByText("The tabs runtime handles keyboard navigation with arrow keys.")).toBeVisible();

  const toastTrigger = page.getByRole("button", { name: "Show toast" });
  await toastTrigger.click();
  await expect(toastTrigger).toBeFocused();
  await expect(page.getByText("Settings saved")).toBeVisible();
  await page.getByRole("button", { name: "Dismiss" }).click();
  await expect(toastTrigger).toBeFocused();
  await expect(page.getByText("Settings saved")).toBeHidden();

  const sheetTrigger = page.getByRole("button", { name: "Open sheet" });
  const sheetSurface = page.getByRole("dialog", { name: "Deployment details" });
  await sheetTrigger.click();
  await expect(sheetSurface).toBeFocused();
  await expect(page.getByText("This panel uses the right edge of the viewport and the same close affordances.")).toBeVisible();
  await page.keyboard.press("Escape");
  await expect(sheetTrigger).toBeFocused();
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
  const swappedDialogTrigger = page.getByRole("button", { name: "Open swapped dialog" });
  const swappedDialog = page.getByRole("dialog", { name: "Swapped fragment" });
  await swappedDialogTrigger.click();
  await expect(swappedDialog).toBeFocused();
  await expect(page.getByText("Delegated runtime hooks mean the client code does not need a manual rebind.")).toBeVisible();
  await page.keyboard.press("Escape");
  await expect(swappedDialogTrigger).toBeFocused();
  await expect(page.getByText("Delegated runtime hooks mean the client code does not need a manual rebind.")).toBeHidden();
});
