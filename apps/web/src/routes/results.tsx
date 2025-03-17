import { useQuery } from '@connectrpc/connect-query';
import { createFileRoute } from '@tanstack/react-router'
import { getVotes } from '../gen/proto/goat/v1/goat-GoatService_connectquery';

export const Route = createFileRoute('/results')({
  component: RouteComponent,
})

function RouteComponent() {

  const data = useQuery(
    getVotes,
    {},
  );

  if(data.isLoading) {
    return <div>Loading...</div>
  }

  return <div className="grid h-screen place-items-center"><div>
    <h1 className='text-4xl'>Is it the goat stack??</h1>
    <div className='text-xl'>Yes: {data.data?.Yes}</div> <div className='text-xl'>No {data.data?.No} </div></div></div>
}
