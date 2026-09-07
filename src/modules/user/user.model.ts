import { t, type UnwrapSchema } from "elysia";

export const UserModel = {
  profileResponse: t.Object({
    username: t.String({ minLength: 3 }),
    email: t.String({ format: "email" }),
  }),

  userNotFound: t.Literal("User not found"),
} as const;

export type UserModel = {
  [k in keyof typeof UserModel]: UnwrapSchema<(typeof UserModel)[k]>;
};
