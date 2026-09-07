import { db } from "../../database";
import { table } from "../../database/schema";
import { eq } from "drizzle-orm";

export const userService = {
  async getUserInfo(userId: string) {
    const userInfo = await db
      .select({
        username: table.user.username,
        email: table.user.email,
      })
      .from(table.user)
      .where(eq(table.user.id, userId))
      .then((res) => res[0]);

    return userInfo;
  },
};
