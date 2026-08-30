import { Elysia } from "elysia";
import { jwt } from "@elysia/jwt";
import { SignInBody, SignUpBody } from "./auth.model";
import { authService } from "./auth.service";

export const authController = new Elysia({ prefix: "/auth" })
  .use(
    jwt({
      name: "jwt",
      secret: process.env.JWT_SECRET!,
      exp: "7d",
    }),
  )
  .post(
    "/sign-up",
    async ({ body }) => {
      return await authService.register(body);
    },
    {
      body: SignUpBody,
    },
  )
  .post(
    "/sign-in",
    async ({ body, jwt, status }) => {
      const result = await authService.login(body);

      if (!result.success || !result.user) {
        return status(401, {
          success: false,
          message: result.message,
        });
      }

      const token = await jwt.sign({
        sub: result.user.id,
      });

      return {
        success: true,
        message: result.message,
        token: token,
        user: result.user,
      };
    },
    {
      body: SignInBody,
    },
  );
