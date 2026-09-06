import { Elysia, t } from "elysia";
import { userService } from "./user.service";

export const userController = new Elysia({ prefix: "/user" }).get(
  "",
  async ({ headers }) => {
    const userId = headers["x-user-id"];
    return await userService.getUserInfo(userId);
  },
  {
    headers: t.Object({
      "x-user-id": t.String({ description: "ID текущего пользователя" }),
    }),
  },
);
