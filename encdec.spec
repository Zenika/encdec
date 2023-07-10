%define debug_package   %{nil}
%define _build_id_links none
%define _name   encdec
%define _prefix /opt
%define _version 0.100
%define _rel 0
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
#Requires: sudo
#Obsoletes: vmman1 > 1.140

%description
Encode and decode AES-256 strings

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
* Mon Jul 10 2023 builder <builder@famillegratton.net> 0.100-0
- new package built with tito

