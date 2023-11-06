%define debug_package   %{nil}
%define _build_id_links none
%define _name   encdec
%define _prefix /opt
%define _version 1.02.00
%define _rel 1
%define _arch x86_64
%define _binaryname encdec

Name:       encdec
Version:    %{_version}
Release:    %{_rel}
Summary:    encdec

Group:      SSL
License:    GPL2.0
URL:        https://github.com/jeanfrancoisgratton/encdec

Source0:    %{name}-%{_version}.tar.gz
BuildArchitectures: x86_64
BuildRequires: gcc

%description
Encode and decode AES-256 strings and files

%prep
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src
PATH=$PATH:/opt/go/bin go build -o %{_sourcedir}/%{_binaryname} .
strip %{_sourcedir}/%{_binaryname}

%clean
rm -rf $RPM_BUILD_ROOT

%pre
exit 0

%install
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}%{_bindir}/%{_binaryname}

%post

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{_binaryname}


%changelog
* Mon Nov 06 2023 RPM Builder <builder@famillegratton.net> 1.02.00-1
- Go version bump (jean-francois@famillegratton.net)

* Mon Nov 06 2023 RPM Builder <builder@famillegratton.net> 1.02.00-0
- Fixed argcount error, version numbering scheme change (jean-
  francois@famillegratton.net)
- Fixed chown issue in packaging (jean-francois@famillegratton.net)

* Wed Aug 02 2023 builder <builder@famillegratton.net> 1.000-1
- Updated changelog and some forgotten relase numbers in packaging scripts
  (jean-francois@famillegratton.net)

* Wed Aug 02 2023 builder <builder@famillegratton.net> 1.000-0
- Prod-ready release: 1.000-0 (jean-francois@famillegratton.net)

* Mon Jul 31 2023 builder <builder@famillegratton.net> 0.200-1
- Doc and version updates (jean-francois@famillegratton.net)

* Mon Jul 31 2023 builder <builder@famillegratton.net> 0.200-0
- Go version bump, also forgotten in previous commits (builder@famillegratton.net)
- Version bump (forgotten in previous commit) (jean-
  francois@famillegratton.net)
- Completed 0.200 (jean-francois@famillegratton.net)
- File Dec-Enc facilities, stub 1 (jean-francois@famillegratton.net)
- New brandh to fully encode/decode files (jean-francois@famillegratton.net)

* Mon Jul 10 2023 builder <builder@famillegratton.net> 0.100-0
- new package built with tito

