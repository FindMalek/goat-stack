import { useMutation } from "@connectrpc/connect-query";
import { createFileRoute, Link, useRouter } from "@tanstack/react-router";
import { vote } from "../gen/proto/goat/v1/goat-GoatService_connectquery";
import { Vote } from "../gen/proto/goat/v1/goat_pb";
import { Button } from "@goat/ui/components/button";

export const Route = createFileRoute("/")({
  component: App,
});

function App() {
  const v = useMutation(vote);
  const { navigate } = useRouter();
  return (
    <div className="text-center">
      <div className="min-h-screen flex flex-col items-center justify-center text-[calc(10px+2vmin)]">
        <p className="text-3xl">Is this the 🐐 stack?</p>

        <div className="my-8 space-x-2">
          <Button
            variant={"outline"}
            disabled={v.isPending}
            onClick={async () => {
              await v.mutateAsync({
                Vote: Vote.YES,
              });
              navigate({ to: "/results" });
            }}
          >
            Yes
          </Button>
          <Button
            variant={"outline"}
            disabled={v.isPending}
            onClick={async () => {
              await v.mutateAsync({
                Vote: Vote.NO,
              });
              navigate({ to: "/results" });
            }}
          >
            No
          </Button>
          <p className="text-sm">
            {" "}
            <Button asChild variant={"link"} size={"sm"}>
              <Link to="/results">View results</Link>
            </Button>
          </p>
        </div>

        <div className="text-lg">
          Because we just ship it at{" "}
          <a
            className="hover:underline"
            href="https://www.openstatus.dev?ref=goatstack.dev"
          >
            OpenStatus
          </a>
          . Here's the goat-stack{" "}
        </div>



        <div>
          <a
            href="https://github.com/thibaultleouay/goat-stack"
            className="text-sm hover:underline"
          >
           GitHub
          </a>
        </div>
      </div>
    </div>
  );
}
