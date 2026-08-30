import { t } from "elysia";
import { createInsertSchema } from "drizzle-typebox";
import { table } from "../../database/schema";

export const createUserSchema = createInsertSchema(table.user, {
  email: t.String({ format: "email" }),
});

export const SignUpBody = t.Omit(createUserSchema, ["id", "salt", "createdAt"]);

export const SignInBody = t.Object({
  email: t.String({ format: "email" }),
  password: t.String(),
});
