################################################################################

# rpmbuilder:relative-pack true

################################################################################

%define  debug_package %{nil}

################################################################################

Summary:         Tiny Redis client for monitor command output top
Name:            redis-monitor-top
Version:         1.2.1
Release:         0%{?dist}
Group:           Applications/System
License:         EKOL
URL:             https://github.com/essentialkaos/redis-monitor-top

Source0:         https://source.kaos.io/%{name}/%{name}-%{version}.tar.bz2

BuildRoot:       %{_tmppath}/%{name}-%{version}-%{release}-root-%(%{__id_u} -n)

BuildRequires:   golang >= 1.8

Provides:        %{name} = %{version}-%{release}

################################################################################

%description
Tiny Redis client for monitor command output top.

################################################################################

%prep
%setup -q

%build
export GOPATH=$(pwd)
go build src/github.com/essentialkaos/%{name}/%{name}.go

%install
rm -rf %{buildroot}

install -dm 755 %{buildroot}%{_bindir}
install -pm 755 %{name} %{buildroot}%{_bindir}/

%clean
rm -rf %{buildroot}

################################################################################

%files
%defattr(-,root,root,-)
%doc LICENSE.EN LICENSE.RU
%{_bindir}/%{name}

################################################################################

%changelog
* Wed Sep 19 2018 Anton Novojilov <andy@essentialkaos.com> - 1.2.1-0
- Code refactoring

* Tue Oct 03 2017 Anton Novojilov <andy@essentialkaos.com> - 1.2.0-0
- Fixed bug with processing custom MONITOR command name
- Improved float output in RPS

* Tue Oct 03 2017 Anton Novojilov <andy@essentialkaos.com> - 1.1.0-0
- Mask non-default MONITOR command name in process tree

* Thu Jul 06 2017 Anton Novojilov <andy@essentialkaos.com> - 1.0.0-0
- Initial build
