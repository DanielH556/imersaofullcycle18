import { revalidateTag } from "next/cache";
import { NextResponse } from "next/server";

export function POST(request: NextResponse, {params}: {params: {eventId: string}}) {
    revalidateTag("events")
    revalidateTag(`events/${params.eventId}`)

    return new Response(null, {status: 204})
}