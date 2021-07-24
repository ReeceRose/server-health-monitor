import AgentInformation from '../components/Tables/AgentInformation';

export default function Home(): JSX.Element {
  return (
    <div className="flex flex-wrap">
      <div className="w-full px-4 mb-12 xl:mb-0">
        <AgentInformation />
      </div>
    </div>
  );
}
