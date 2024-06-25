/**
 * v0 by Vercel.
 * @see https://v0.dev/t/m2e4Ll9NicL
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */
"use client";

import { useState, useMemo, Fragment } from "react";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from "@/components/ui/select";
import {
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
} from "@/components/ui/table";
import { Pagination } from "@/components/ui/pagination";

interface Meeting {
  id: number;
  date: Date;
  location: string;
  role: string;
  description: string;
}

export default function Component() {
  const [filter, setFilter] = useState<string>("today");
  const [pageSize, setPageSize] = useState<number>(25);
  const [currentPage, setCurrentPage] = useState<number>(1);
  const meetings = useMemo<Meeting[]>(() => {
    const today = new Date();
    const oneWeekFromNow = new Date(today.getTime() + 7 * 24 * 60 * 60 * 1000);
    const oneMonthFromNow = new Date(
      today.getFullYear(),
      today.getMonth() + 1,
      1
    );
    return [
      {
        id: 1,
        date: new Date(
          today.getFullYear(),
          today.getMonth(),
          today.getDate(),
          10,
          0
        ),
        location: "Conference Room A",
        role: "Presenter",
        description: "Quarterly Business Review",
      },
      {
        id: 2,
        date: new Date(
          today.getFullYear(),
          today.getMonth(),
          today.getDate(),
          14,
          0
        ),
        location: "Virtual",
        role: "Attendee",
        description: "Product Strategy Meeting",
      },
      {
        id: 3,
        date: new Date(
          today.getFullYear(),
          today.getMonth(),
          today.getDate() + 2,
          9,
          30
        ),
        location: "Conference Room B",
        role: "Organizer",
        description: "Project Kickoff",
      },
      {
        id: 4,
        date: new Date(
          today.getFullYear(),
          today.getMonth(),
          today.getDate() + 4,
          16,
          0
        ),
        location: "Virtual",
        role: "Attendee",
        description: "Team Sync",
      },
      {
        id: 5,
        date: new Date(
          today.getFullYear(),
          today.getMonth(),
          today.getDate() + 7,
          11,
          0
        ),
        location: "Conference Room C",
        role: "Presenter",
        description: "Client Presentation",
      },
      {
        id: 6,
        date: new Date(
          today.getFullYear(),
          today.getMonth(),
          today.getDate() + 14,
          15,
          0
        ),
        location: "Virtual",
        role: "Attendee",
        description: "Training Session",
      },
      {
        id: 7,
        date: new Date(
          today.getFullYear(),
          today.getMonth(),
          today.getDate() + 21,
          10,
          30
        ),
        location: "Conference Room A",
        role: "Organizer",
        description: "Team Building Activity",
      },
    ].filter((meeting) => {
      if (filter === "today") {
        return meeting.date.toDateString() === today.toDateString();
      } else if (filter === "week") {
        return meeting.date >= today && meeting.date < oneWeekFromNow;
      } else if (filter === "next7days") {
        return meeting.date >= today && meeting.date < oneWeekFromNow;
      } else if (filter === "month") {
        return meeting.date >= today && meeting.date < oneMonthFromNow;
      }
    });
  }, [filter]);

  const totalPages = Math.ceil(meetings.length / pageSize);
  const startIndex = (currentPage - 1) * pageSize;
  const endIndex = startIndex + pageSize;
  const visibleMeetings = meetings.slice(startIndex, endIndex);

  const groupedMeetings = visibleMeetings.reduce<{ [key: string]: Meeting[] }>(
    (acc, meeting) => {
      const date = meeting.date.toDateString();
      if (!acc[date]) {
        acc[date] = [];
      }
      acc[date].push(meeting);
      return acc;
    },
    {}
  );

  return (
    <div className="p-4 md:p-8">
      <div className="flex items-center justify-between mb-4">
        <div className="flex items-center gap-2">
          <Label htmlFor="filter">Filter:</Label>
          <Select value={filter} onValueChange={setFilter}>
            <SelectTrigger>
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="today">Today</SelectItem>
              <SelectItem value="week">This Week</SelectItem>
              <SelectItem value="next7days">Next 7 Days</SelectItem>
              <SelectItem value="month">This Month</SelectItem>
            </SelectContent>
          </Select>
        </div>
        <div className="flex items-center gap-2">
          <Label htmlFor="pageSize">Show:</Label>
          <Select
            value={pageSize.toString()}
            onValueChange={(value) => setPageSize(parseInt(value))}
          >
            <SelectTrigger>
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="25">25</SelectItem>
              <SelectItem value="50">50</SelectItem>
              <SelectItem value="100">100</SelectItem>
              <SelectItem value="250">250</SelectItem>
            </SelectContent>
          </Select>
        </div>
      </div>
      <div className="overflow-x-auto">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Date</TableHead>
              <TableHead>Time</TableHead>
              <TableHead>Location</TableHead>
              <TableHead>Role</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {Object.entries(groupedMeetings).map(([date, meetings]) => (
              <Fragment key={date}>
                <TableRow>
                  <TableCell colSpan={5} className="font-semibold">
                    {date}
                  </TableCell>
                </TableRow>
                {meetings.map((meeting) => (
                  <TableRow key={meeting.id}>
                    <TableCell>
                      {meeting.date.toLocaleTimeString([], {
                        hour: "2-digit",
                        minute: "2-digit",
                      })}
                    </TableCell>
                    <TableCell>{meeting.location}</TableCell>
                    <TableCell>{meeting.role}</TableCell>
                    <TableCell>{meeting.description}</TableCell>
                  </TableRow>
                ))}
              </Fragment>
            ))}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
