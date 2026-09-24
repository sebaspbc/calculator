import { describe, it, expect, vi, afterEach } from "vitest";
import { calculate } from "./api";

describe("calculate", () => {
  afterEach(() => {
    vi.unstubAllGlobals();
  });

  it("returns the result on success", async () => {
    vi.stubGlobal("fetch", vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ result: 8 }),
    }));

    const result = await calculate({ operation: "add", a: 5, b: 3 });
    expect(result).toBe(8);
  });

  it("throws with the backend error message on failure", async () => {
    vi.stubGlobal("fetch", vi.fn().mockResolvedValue({
      ok: false,
      json: async () => ({ error: "division by zero is not allowed" }),
    }));

    await expect(calculate({ operation: "divide", a: 1, b: 0 }))
      .rejects.toThrow("division by zero is not allowed");
  });
});