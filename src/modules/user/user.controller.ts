import { Elysia } from "elysia";

export const userController = new Elysia({ prefix: "/user" }).get(
  "/",
  async () => {
    return "Данные пользователя";
  },
);
