import { useForm, Controller } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import {
  Popover,
  PopoverTrigger,
  PopoverContent,
} from "@/components/ui/popover";
import { Calendar } from "@/components/ui/calendar";
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from "@/components/ui/select";
import { Checkbox } from "@/components/ui/checkbox";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";

// Define validation schema using zod
const formSchema = z.object({
  purpose: z.string().min(1, "Purpose is required"),
  date: z.string().min(1, "Date is required"),
  startTime: z.string().min(1, "Start Time is required"),
  endTime: z.string().min(1, "End Time is required"),
  agendaItem: z.string().optional(),
  presenter: z.string().optional(),
  duration: z.string().optional(),
  info: z.boolean().optional(),
  decision: z.boolean().optional(),
});

type FormSchema = z.infer<typeof formSchema>;

export default function NewMeetingForm() {
  const {
    register,
    handleSubmit,
    control,
    formState: { errors },
  } = useForm<FormSchema>({
    resolver: zodResolver(formSchema),
  });

  const onSubmit = (data: FormSchema) => {
    console.log(data);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="p-4 md:p-6">
      <div className="grid grid-cols-[1fr_150px] gap-6">
        <div className="grid gap-6">
          <Card>
            <CardHeader>
              <CardTitle>Meeting Templates</CardTitle>
              <CardDescription>
                Select a template to get started or create a new meeting.
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="grid grid-cols-2 gap-4">
                <Button variant="outline" className="col-span-1">
                  <ClipboardListIcon className="w-4 h-4 mr-2" />
                  Weekly Team Meeting
                </Button>
                <Button variant="outline" className="col-span-1">
                  <BarChartIcon className="w-4 h-4 mr-2" />
                  Monthly Project Review
                </Button>
                <Button variant="outline" className="col-span-1">
                  <PresentationIcon className="w-4 h-4 mr-2" />
                  Product Launch
                </Button>
                <Button variant="outline" className="col-span-1">
                  <PlusIcon className="w-4 h-4 mr-2" />
                  New Meeting
                </Button>
              </div>
            </CardContent>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Meeting Details</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="grid gap-4">
                <div>
                  <Label htmlFor="purpose">Purpose</Label>
                  <Input
                    id="purpose"
                    placeholder="Enter meeting purpose"
                    {...register("purpose")}
                  />
                  {errors.purpose && (
                    <span className="text-red-500">
                      {errors.purpose.message}
                    </span>
                  )}
                </div>
                <div>
                  <Label>Date and Time</Label>
                  <div className="grid grid-cols-2 gap-2">
                    <div>
                      <Label htmlFor="date">Date</Label>
                      <Popover>
                        <PopoverTrigger asChild>
                          <Button
                            variant="outline"
                            className="w-full justify-start text-left font-normal"
                          >
                            <CalendarDaysIcon className="mr-1 h-4 w-4 -translate-x-1" />
                            Pick a date
                          </Button>
                        </PopoverTrigger>
                        <PopoverContent className="w-auto p-0" align="start">
                          <Calendar mode="single" initialFocus />
                        </PopoverContent>
                      </Popover>
                    </div>
                    <div>
                      <Label htmlFor="startTime">Start Time</Label>
                      <Input
                        id="startTime"
                        type="time"
                        {...register("startTime")}
                      />
                      {errors.startTime && (
                        <span className="text-red-500">
                          {errors.startTime.message}
                        </span>
                      )}
                    </div>
                    <div>
                      <Label htmlFor="endTime">End Time</Label>
                      <Input
                        id="endTime"
                        type="time"
                        {...register("endTime")}
                      />
                      {errors.endTime && (
                        <span className="text-red-500">
                          {errors.endTime.message}
                        </span>
                      )}
                    </div>
                  </div>
                  <div>
                    <Label>Suggested Times</Label>
                    <div className="grid grid-cols-3 gap-2">
                      <Button variant="outline" className="col-span-1">
                        10:00 AM - 11:30 AM
                      </Button>
                      <Button variant="outline" className="col-span-1">
                        2:00 PM - 3:30 PM
                      </Button>
                      <Button variant="outline" className="col-span-1">
                        4:00 PM - 5:30 PM
                      </Button>
                    </div>
                  </div>
                </div>
                <div>
                  <Label>Agenda</Label>
                  <div className="grid gap-2">
                    <div className="grid grid-cols-5 gap-2">
                      <Input
                        placeholder="Agenda item"
                        className="col-span-3"
                        {...register("agendaItem")}
                      />
                      <Controller
                        name="presenter"
                        control={control}
                        render={({ field }) => (
                          <Select {...field}>
                            <SelectTrigger>
                              <SelectValue placeholder="Presenter" />
                            </SelectTrigger>
                            <SelectContent>
                              <SelectItem value="presenter">
                                Presenter
                              </SelectItem>
                              <SelectItem value="john">John Doe</SelectItem>
                              <SelectItem value="jane">Jane Smith</SelectItem>
                            </SelectContent>
                          </Select>
                        )}
                      />
                      <Input
                        type="time"
                        placeholder="Duration"
                        className="col-span-1"
                        {...register("duration")}
                      />
                    </div>
                    <div className="flex items-center gap-2">
                      <Controller
                        name="info"
                        control={control}
                        render={({ field }) => (
                          <>
                            <Checkbox
                              checked={field.value}
                              onCheckedChange={field.onChange}
                              id="info"
                            />
                            <Label htmlFor="info">For Information</Label>
                          </>
                        )}
                      />
                      <Controller
                        name="decision"
                        control={control}
                        render={({ field }) => (
                          <>
                            <Checkbox
                              checked={field.value}
                              onCheckedChange={field.onChange}
                              id="decision"
                            />
                            <Label htmlFor="decision">For Decision</Label>
                          </>
                        )}
                      />
                    </div>
                  </div>
                  <Button variant="outline" className="mt-2">
                    Add Agenda Item
                  </Button>
                </div>
              </div>
            </CardContent>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Room Selection</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="grid gap-4">
                <div className="flex items-center gap-2">
                  {/* <Controller
                    name="site"
                    control={control}
                    render={({ field }) => (
                      <Select >
                        <SelectTrigger>
                          <SelectValue placeholder="Site" />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectItem value="site">Site</SelectItem>
                          <SelectItem value="site1">Site 1</SelectItem>
                          <SelectItem value="site2">Site 2</SelectItem>
                        </SelectContent>
                      </Select>
                    )}
                  />
                  <Controller
                    name="building"
                    control={control}
                    render={({ field }) => (
                      <Select {...field}>
                        <SelectTrigger>
                          <SelectValue placeholder="Building" />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectItem value="building">Building</SelectItem>
                          <SelectItem value="building1">Building 1</SelectItem>
                          <SelectItem value="building2">Building 2</SelectItem>
                        </SelectContent>
                      </Select>
                    )}
                  />
                  <Controller
                    name="floor"
                    control={control}
                    render={({ field }) => (
                      <Select {...field}>
                        <SelectTrigger>
                          <SelectValue placeholder="Floor" />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectItem value="floor">Floor</SelectItem>
                          <SelectItem value="floor1">Floor 1</SelectItem>
                          <SelectItem value="floor2">Floor 2</SelectItem>
                        </SelectContent>
                      </Select>
                    )}
                  /> */}
                </div>
                <div className="grid grid-cols-2 gap-4">
                  <Card className="relative">
                    <CardHeader>
                      <CardTitle>Room 101</CardTitle>
                      <CardDescription>Available</CardDescription>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm">Capacity: 10 people</p>
                      <p className="text-sm">
                        Amenities: Projector, Whiteboard
                      </p>
                    </CardContent>
                    <CardFooter>
                      <Button>Select</Button>
                    </CardFooter>
                  </Card>
                  <Card className="relative">
                    <CardHeader>
                      <CardTitle>Room 202</CardTitle>
                      <CardDescription>Booked</CardDescription>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm">Capacity: 15 people</p>
                      <p className="text-sm">
                        Amenities: TV, Videoconferencing
                      </p>
                      <p className="text-sm text-gray-500 dark:text-gray-400">
                        Booked by: Jane Smith
                      </p>
                    </CardContent>
                    <CardFooter>
                      <Button variant="outline" disabled>
                        Select
                      </Button>
                    </CardFooter>
                  </Card>
                </div>
              </div>
            </CardContent>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Attendees</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="grid gap-4">
                <div className="flex items-center gap-2">
                  <Input placeholder="Add attendees" />
                  <Button>Invite</Button>
                </div>
                <div className="grid grid-cols-3 gap-4">
                  <div className="flex items-center gap-2">
                    <Avatar>
                      <AvatarImage src="/placeholder-user.jpg" />
                      <AvatarFallback>JD</AvatarFallback>
                    </Avatar>
                    <div>
                      <p className="font-medium">John Doe</p>
                      <p className="text-sm text-gray-500 dark:text-gray-400">
                        Accepted
                      </p>
                    </div>
                  </div>
                  <div className="flex items-center gap-2">
                    <Avatar>
                      <AvatarImage src="/placeholder-user.jpg" />
                      <AvatarFallback>JS</AvatarFallback>
                    </Avatar>
                    <div>
                      <p className="font-medium">Jane Smith</p>
                      <p className="text-sm text-gray-500 dark:text-gray-400">
                        Remote
                      </p>
                    </div>
                  </div>
                  <div className="flex items-center gap-2">
                    <Avatar>
                      <AvatarImage src="/placeholder-user.jpg" />
                      <AvatarFallback>BJ</AvatarFallback>
                    </Avatar>
                    <div>
                      <p className="font-medium">Bob Johnson</p>
                      <p className="text-sm text-gray-500 dark:text-gray-400">
                        Room 101
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Additional Services</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="flex items-center gap-2">
                <Button variant="outline">
                  <CoffeeIcon className="w-4 h-4 mr-2" />
                  Order Catering
                </Button>
                <Button variant="outline">
                  <PresentationIcon className="w-4 h-4 mr-2" />
                  Request Presentation Setup
                </Button>
              </div>
            </CardContent>
          </Card>
        </div>
        <div className="sticky top-0 self-start">
          <nav className="flex flex-col gap-2">
            <a
              href="#meeting-templates"
              className="px-2 py-1 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800"
            >
              Meeting Templates
            </a>
            <a
              href="#meeting-details"
              className="px-2 py-1 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800"
            >
              Meeting Details
            </a>
            <a
              href="#room-selection"
              className="px-2 py-1 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800"
            >
              Room Selection
            </a>
            <a
              href="#attendees"
              className="px-2 py-1 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800"
            >
              Attendees
            </a>
            <a
              href="#additional-services"
              className="px-2 py-1 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800"
            >
              Additional Services
            </a>
          </nav>
        </div>
      </div>
    </form>
  );
}

function BarChartIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <line x1="12" x2="12" y1="20" y2="10" />
      <line x1="18" x2="18" y1="20" y2="4" />
      <line x1="6" x2="6" y1="20" y2="16" />
    </svg>
  );
}

function CalendarDaysIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M8 2v4" />
      <path d="M16 2v4" />
      <rect width="18" height="18" x="3" y="4" rx="2" />
      <path d="M3 10h18" />
      <path d="M8 14h.01" />
      <path d="M12 14h.01" />
      <path d="M16 14h.01" />
      <path d="M8 18h.01" />
      <path d="M12 18h.01" />
      <path d="M16 18h.01" />
    </svg>
  );
}

function ClipboardListIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <rect width="8" height="4" x="8" y="2" rx="1" ry="1" />
      <path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2" />
      <path d="M12 11h4" />
      <path d="M12 16h4" />
      <path d="M8 11h.01" />
      <path d="M8 16h.01" />
    </svg>
  );
}

function CoffeeIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M10 2v2" />
      <path d="M14 2v2" />
      <path d="M16 8a1 1 0 0 1 1 1v8a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4V9a1 1 0 0 1 1-1h14a4 4 0 1 1 0 8h-1" />
      <path d="M6 2v2" />
    </svg>
  );
}

function PlusIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M5 12h14" />
      <path d="M12 5v14" />
    </svg>
  );
}

function PresentationIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M2 3h20" />
      <path d="M21 3v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V3" />
      <path d="m7 21 5-5 5 5" />
    </svg>
  );
}
