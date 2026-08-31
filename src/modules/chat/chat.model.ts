import { createInsertSchema } from "drizzle-typebox";
import { table } from "../../database/schema";
import { t } from "elysia";

export const createChatSchema = createInsertSchema(table.chat);

export const CreateChatBody = t.Omit(createChatSchema, ["id", "createdAt"]);

export const AddParticipantBody = t.Object({
  userId: t.String({ description: "ID пользователя, которого добавляем" }),
});
