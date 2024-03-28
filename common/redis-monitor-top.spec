################################################################################

%define  debug_package %{nil}

################################################################################

Summary:         Tiny Redis client for monitor command output top
Name:            redis-monitor-top
Version:         1.3.3
Release:         0%{?dist}
Group:           Applications/System
License:         Apache License, Version 2.0
URL:             https://kaos.sh/redis-monitor-top

Source0:         https://source.kaos.st/%{name}/%{name}-%{version}.tar.bz2

BuildRoot:       %{_tmppath}/%{name}-%{version}-%{release}-root-%(%{__id_u} -n)

BuildRequires:   golang >= 1.21

Provides:        %{name} = %{version}-%{release}

################################################################################

%description
Tiny Redis client for monitor command output top.

################################################################################

%prep
%setup -q

%build
if [[ ! -d "%{name}/vendor" ]] ; then
  echo "This package requires vendored dependencies"
  exit 1
fi

pushd %{name}
  go build %{name}.go
  cp LICENSE ..
popd

%install
rm -rf %{buildroot}

install -dm 755 %{buildroot}%{_bindir}
install -pm 755 %{name}/%{name} %{buildroot}%{_bindir}/

%clean
rm -rf %{buildroot}

%post
if [[ -d %{_sysconfdir}/bash_completion.d ]] ; then
  %{name} --completion=bash 1> %{_sysconfdir}/bash_completion.d/%{name} 2>/dev/null
fi

if [[ -d %{_datarootdir}/fish/vendor_completions.d ]] ; then
  %{name} --completion=fish 1> %{_datarootdir}/fish/vendor_completions.d/%{name}.fish 2>/dev/null
fi

if [[ -d %{_datadir}/zsh/site-functions ]] ; then
  %{name} --completion=zsh 1> %{_datadir}/zsh/site-functions/_%{name} 2>/dev/null
fi

%postun
if [[ $1 == 0 ]] ; then
  if [[ -f %{_sysconfdir}/bash_completion.d/%{name} ]] ; then
    rm -f %{_sysconfdir}/bash_completion.d/%{name} &>/dev/null || :
  fi

  if [[ -f %{_datarootdir}/fish/vendor_completions.d/%{name}.fish ]] ; then
    rm -f %{_datarootdir}/fish/vendor_completions.d/%{name}.fish &>/dev/null || :
  fi

  if [[ -f %{_datadir}/zsh/site-functions/_%{name} ]] ; then
    rm -f %{_datadir}/zsh/site-functions/_%{name} &>/dev/null || :
  fi
fi

################################################################################

%files
%defattr(-,root,root,-)
%doc LICENSE
%{_bindir}/%{name}

################################################################################

%changelog
* Thu Mar 28 2024 Anton Novojilov <andy@essentialkaos.com> - 1.3.3-0
- Improved support information gathering
- Code refactoring
- Dependencies update

* Thu Dec 01 2022 Anton Novojilov <andy@essentialkaos.com> - 1.3.2-1
- Fixed build using sources from source.kaos.st

* Tue Mar 29 2022 Anton Novojilov <andy@essentialkaos.com> - 1.3.2-0
- Package ek updated to the latest stable version
- Removed pkg.re usage
- Added module info
- Added Dependabot configuration

* Thu Oct 17 2019 Anton Novojilov <andy@essentialkaos.com> - 1.3.1-0
- ek package updated to the latest stable version

* Sat Jun 15 2019 Anton Novojilov <andy@essentialkaos.com> - 1.3.0-0
- ek package updated to the latest stable version
- Added completion generation for bash, zsh and fish

* Sat Oct 20 2018 Anton Novojilov <andy@essentialkaos.com> - 1.2.2-0
- Show usage info if '-h' passed without any value

* Wed Sep 19 2018 Anton Novojilov <andy@essentialkaos.com> - 1.2.1-0
- Code refactoring

* Tue Oct 03 2017 Anton Novojilov <andy@essentialkaos.com> - 1.2.0-0
- Fixed bug with processing custom MONITOR command name
- Improved float output in RPS

* Tue Oct 03 2017 Anton Novojilov <andy@essentialkaos.com> - 1.1.0-0
- Mask non-default MONITOR command name in process tree

* Thu Jul 06 2017 Anton Novojilov <andy@essentialkaos.com> - 1.0.0-0
- Initial build
