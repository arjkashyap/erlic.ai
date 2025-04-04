import { NextResponse } from "next/server";
import { cookies } from "next/headers";
import { checkAuth } from "@/lib/auth-utils";

export async function GET() {
  const cookieStore = await cookies();
  const sessionCookie = cookieStore.get("_gothic_session");
  const authResult = await checkAuth(sessionCookie);
  
  return NextResponse.json(authResult);
} 