"use client";

import { useRouter } from "next/navigation";
import { APPNAME } from "../global";
import NewMeetingForm from "@/components/meeting_new";
import EventsOverview from "@/components/events_overview";

export default function Page() {
  const router = useRouter();
  return (
    <div className="space-x-2 h-[90vh]">
      <EventsOverview />
    </div>
  );
}
