import EventCard from "./components/EventCard";
import { Title } from "./components/Title";
import { EventModel } from "./models";

// props.event.date

export async function getEvents(): Promise<EventModel[]> {
  const response = await fetch("http://localhost:8080/events", {
    cache: "no-store"
  });

  return (await response.json()).events;
}

export default async function Home() {
  const events = await getEvents();

  return (
    <main>
      <Title>Eventos Disponíveis</Title>
      <div className="mt-8 sm:grid sm:grid-cols-auto-fit-cards flex flex-wrap justify-center gap-x-2 gap-y-4">
        {events.map((event) => (
          <EventCard key={event.id} event={event} />
        ))}
      </div>
    </main>
  );
}
