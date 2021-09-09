import Link from 'next/link';

export default function Navbar(): JSX.Element {
  return (
    <nav className="absolute top-0 left-0 z-10 flex items-center w-full p-4 bg-transparent md:flex-row md:flex-nowrap md:justify-start">
      <div className="flex flex-wrap items-center justify-between w-full px-4 mx-autp md:flex-nowrap md:px-10">
        <Link href="/">
          <a className="hidden text-sm font-semibold text-white uppercase lg:inline-block">
            Dashboard
          </a>
        </Link>

        <ul className="flex-col items-center hidden list-none md:flex-row md:flex">
          <button
            className="block px-3 py-1 text-gray-200"
            // onClick={(e) => e.preventDefault()}
          >
            <i className="fas fa-bell"></i>
          </button>
        </ul>
      </div>
    </nav>
  );
}
