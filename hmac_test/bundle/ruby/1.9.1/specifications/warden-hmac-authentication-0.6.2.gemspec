# -*- encoding: utf-8 -*-
# stub: warden-hmac-authentication 0.6.2 ruby lib

Gem::Specification.new do |s|
  s.name = "warden-hmac-authentication"
  s.version = "0.6.2"

  s.required_rubygems_version = Gem::Requirement.new(">= 0") if s.respond_to? :required_rubygems_version=
  s.authors = ["Felix Gilcher", "Florian Gilcher"]
  s.date = "2012-07-09"
  s.description = "This gem provides request authentication via [HMAC](http://en.wikipedia.org/wiki/Hmac). The main usage is request based, noninteractive\n  authentication for API implementations. Two strategies are supported that differ mainly in how the authentication information is\n  transferred to the server: One header-based authentication method and one query-based. The authentication scheme is in some parts based\n  on ideas laid out in this article and the following discussion: \n  http://broadcast.oreilly.com/2009/12/principles-for-standardized-rest-authentication.html\n\n  The gem also provides a small helper class that can be used to generate request signatures."
  s.email = ["felix.gilcher@asquera.de", "florian.gilcher@asquera.de"]
  s.executables = ["warden-hmac-authentication"]
  s.files = ["bin/warden-hmac-authentication"]
  s.homepage = "https://github.com/Asquera/warden-hmac-authentication"
  s.require_paths = ["lib"]
  s.rubygems_version = "2.1.9"
  s.summary = "Provides request based, non-interactive authentication for APIs"

  if s.respond_to? :specification_version then
    s.specification_version = 3

    if Gem::Version.new(Gem::VERSION) >= Gem::Version.new('1.2.0') then
      s.add_runtime_dependency(%q<addressable>, [">= 0"])
      s.add_runtime_dependency(%q<rack>, [">= 0"])
      s.add_runtime_dependency(%q<warden>, [">= 0"])
      s.add_development_dependency(%q<rake>, [">= 0"])
      s.add_development_dependency(%q<rack-test>, [">= 0"])
      s.add_development_dependency(%q<riot>, [">= 0"])
      s.add_development_dependency(%q<timecop>, [">= 0"])
      s.add_development_dependency(%q<simplecov>, [">= 0"])
      s.add_development_dependency(%q<simplecov-html>, [">= 0"])
      s.add_development_dependency(%q<trollop>, [">= 0"])
    else
      s.add_dependency(%q<addressable>, [">= 0"])
      s.add_dependency(%q<rack>, [">= 0"])
      s.add_dependency(%q<warden>, [">= 0"])
      s.add_dependency(%q<rake>, [">= 0"])
      s.add_dependency(%q<rack-test>, [">= 0"])
      s.add_dependency(%q<riot>, [">= 0"])
      s.add_dependency(%q<timecop>, [">= 0"])
      s.add_dependency(%q<simplecov>, [">= 0"])
      s.add_dependency(%q<simplecov-html>, [">= 0"])
      s.add_dependency(%q<trollop>, [">= 0"])
    end
  else
    s.add_dependency(%q<addressable>, [">= 0"])
    s.add_dependency(%q<rack>, [">= 0"])
    s.add_dependency(%q<warden>, [">= 0"])
    s.add_dependency(%q<rake>, [">= 0"])
    s.add_dependency(%q<rack-test>, [">= 0"])
    s.add_dependency(%q<riot>, [">= 0"])
    s.add_dependency(%q<timecop>, [">= 0"])
    s.add_dependency(%q<simplecov>, [">= 0"])
    s.add_dependency(%q<simplecov-html>, [">= 0"])
    s.add_dependency(%q<trollop>, [">= 0"])
  end
end
