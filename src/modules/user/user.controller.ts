import { Elysia, t } from "elysia";
import { userService } from "./user.service";
import { UserModel } from "./user.model";

export const userController = new Elysia({ prefix: "/user" }).get(
  "",
  async ({ headers, status }) => {
    const userId = headers["x-user-id"];
    const user = await userService.getUserInfo(userId);

    if (!user) {
      return status(404, "User not found");
    }

    return user;
  },
  {
    headers: t.Object({
      "x-user-id": t.String({ description: "ID текущего пользователя" }),
    }),
    response: {
      200: UserModel.profileResponse,
      404: UserModel.userNotFound,
    },
  },
);
