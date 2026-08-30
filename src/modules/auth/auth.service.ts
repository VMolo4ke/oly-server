import { eq } from "drizzle-orm";
import { db } from "../../database/index";
import { table } from "../../database/schema";
import { SignInBody, SignUpBody } from "./auth.model";

export const authService = {
  async register(data: typeof SignUpBody.static) {
    const hashedPassword = await Bun.password.hash(data.password, {
      algorithm: "bcrypt",
      cost: 10,
    });

    const [newUser] = await db
      .insert(table.user)
      .values({
        email: data.email,
        username: data.username,
        password: hashedPassword,
      })
      .returning();

    return {
      success: true,
      message: "Пользователь успешно зарегистрирован",
      user: { id: newUser.id, email: newUser.email },
    };
  },

  async login(data: typeof SignInBody.static) {
    const [user] = await db
      .select()
      .from(table.user)
      .where(eq(table.user.email, data.email))
      .limit(1);

    if (!user) {
      return { success: false, message: "Неверный email или password" };
    }

    const isPasswordValid = await Bun.password.verify(
      data.password,
      user.password,
    );

    if (!isPasswordValid) {
      return { success: false, message: "Неверный email или password" };
    }

    return {
      success: true,
      message: "Вы успешно вошли в систему",
      user: {
        id: user.id,
        username: user.username,
        email: user.email,
      },
    };
  },
};
