import { expect, test } from "@playwright/test";

test("form workflow example exposes a realistic release request form", async ({ page }) => {
  await page.goto("/examples/form-workflow");

  await expect(page.getByRole("heading", { name: "Testable form workflow" })).toBeVisible();
  await expect(page.getByText("Review before deploy")).toBeVisible();
  await expect(page.getByLabel("Service name")).toHaveValue("search-api");
  await expect(page.getByLabel(/Change ticket/)).toHaveValue("CHG-1842");
  await expect(page.getByText("Ticket ownership has not been assigned yet.")).toBeVisible();
  await expect(page.getByRole("combobox")).toHaveValue("eu-morning");
});

test("runtime workbench example keeps interactive controls testable", async ({ page }) => {
  await page.goto("/examples/runtime-workbench");

  await page.getByRole("button", { name: "Open approval dialog" }).click();
  await expect(page.getByText("Confirm the canary has remained healthy for five minutes before promoting the release.")).toBeVisible();
  await page.keyboard.press("Escape");
  await expect(page.getByText("Confirm the canary has remained healthy for five minutes before promoting the release.")).toBeHidden();

  await page.getByRole("button", { name: "Quick actions" }).click();
  await expect(page.getByText("Pause rollout")).toBeVisible();
  await page.mouse.click(5, 5);
  await expect(page.getByText("Pause rollout")).toBeHidden();

  await page.getByRole("tab", { name: "Overview" }).focus();
  await page.keyboard.press("ArrowRight");
  await expect(page.getByRole("tab", { name: "Activity" })).toBeFocused();
  await expect(page.getByText("Keyboard arrow keys move focus between tabs without a client-side framework.")).toBeVisible();

  await page.getByRole("button", { name: "Show save toast" }).click();
  await expect(page.getByText("Incident note saved")).toBeVisible();

  await page.getByRole("button", { name: "Open operator sheet" }).click();
  await expect(page.getByText("The sheet is useful for secondary details that should not crowd the main grid.")).toBeVisible();
  await page.keyboard.press("Escape");
  await expect(page.getByText("The sheet is useful for secondary details that should not crowd the main grid.")).toBeHidden();
});
