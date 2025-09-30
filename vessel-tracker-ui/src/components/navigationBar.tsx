import { Navbar, NavbarBrand, NavbarContent, NavbarItem } from "@heroui/navbar";
import { Link } from "@heroui/link";
import { Image } from "@heroui/image";

export default function NavigationBar() {
  return (
    <Navbar>
      <NavbarBrand justify="start" className="flex gap-2">
        <Image
          src="/vessel-tracker-logo.png"
          alt="Vessel Tracker Logo"
          width={70}
          height={70}
        />
        <p className="font-bold text-inherit">Vessel-Tracker </p>
        <p className="text-inherit text-sm"> by dh1542</p>
      </NavbarBrand>
      <NavbarContent className="hidden sm:flex gap-7" justify="center">
        <NavbarItem>
          <Link href="#">Options</Link>
        </NavbarItem>
        <NavbarItem>
          <Link aria-current="page" href="#">
            Map Layers
          </Link>
        </NavbarItem>
      </NavbarContent>
      <NavbarContent justify="end">
        <NavbarItem className="hidden lg:flex ">
          <Link isExternal href="https://github.com/dh1542/vessel-tracker">
            <Image
              src="/github-mark.png"
              alt="Github Logo"
              width={20}
              height={20}
            />
            <span>&nbsp;</span>
            <p className="text-black"> by dh1542</p>
          </Link>
        </NavbarItem>
      </NavbarContent>
    </Navbar>
  );
}
