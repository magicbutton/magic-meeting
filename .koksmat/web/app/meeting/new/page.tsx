"use client";

import { useRouter } from "next/navigation";

import NewMeetingForm from "@/components/meeting_new";

export default function Page() {
  const router = useRouter();
  return (
    <div className="space-x-2 h-[90vh]">
      <NewMeetingForm />
    </div>
  );
}
