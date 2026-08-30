import { pgTable, varchar, timestamp } from "drizzle-orm/pg-core";
import { createId } from "@paralleldrive/cuid2";

export const user = pgTable("user", {
  id: varchar("id")
    .$default(() => createId())
    .primaryKey(),
  username: varchar("username", { length: 50 }).notNull().unique(),
  password: varchar("password", { length: 255 }).notNull(),
  email: varchar("email", { length: 255 }).notNull().unique(),
  createdAt: timestamp("created_at").defaultNow().notNull(),
});

export const chat = pgTable("chat", {
  id: varchar("id")
    .$default(() => createId())
    .primaryKey(),
  creatorId: varchar("creator_id").references(() => user.id, {
    onDelete: "set null",
  }),
  createdAt: timestamp("created_at").defaultNow().notNull(),
});

export const message = pgTable("message", {
  id: varchar("id")
    .$default(() => createId())
    .primaryKey(),

  text: varchar("text", { length: 4000 }).notNull(),

  chatId: varchar("chat_id")
    .notNull()
    .references(() => chat.id, { onDelete: "cascade" }),

  senderId: varchar("sender_id")
    .notNull()
    .references(() => user.id, { onDelete: "cascade" }),

  createdAt: timestamp("created_at").defaultNow().notNull(),
});

export const voiceCall = pgTable("voice_call", {
  id: varchar("id")
    .$default(() => createId())
    .primaryKey(),
  chatId: varchar("chat_id")
    .notNull()
    .references(() => chat.id, { onDelete: "cascade" }),
  status: varchar("status", { length: 20 }).default("active").notNull(),
  createdAt: timestamp("created_at").defaultNow().notNull(),
});

export const voiceParticipant = pgTable("voice_participant", {
  id: varchar("id")
    .$default(() => createId())
    .primaryKey(),
  callId: varchar("call_id")
    .notNull()
    .references(() => voiceCall.id, { onDelete: "cascade" }),
  userId: varchar("user_id")
    .notNull()
    .references(() => user.id, { onDelete: "cascade" }),
  joinedAt: timestamp("joined_at").defaultNow().notNull(),
});

export const table = {
  user,
  chat,
  message,
  voiceCall,
  voiceParticipant,
} as const;

export type Table = typeof table;
