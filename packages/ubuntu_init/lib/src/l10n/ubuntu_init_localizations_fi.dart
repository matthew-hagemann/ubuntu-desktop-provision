import 'ubuntu_init_localizations.dart';

/// The translations for Finnish (`fi`).
class UbuntuInitLocalizationsFi extends UbuntuInitLocalizations {
  UbuntuInitLocalizationsFi([String locale = 'fi']) : super(locale);

  @override
  String get welcomePageTitle => 'Tervetuloa';

  @override
  String welcomePageHeader(Object distro) {
    return 'Tervetuloa, tämä on $distro';
  }

  @override
  String get welcomeWhatsNew => 'MITÄ UUTTA?';

  @override
  String get welcomeStartTitle => 'Suoraviivaistettu aloitus';

  @override
  String get welcomeStartSubtitle => 'Aiempaa siistimpi asennuskokemus.';

  @override
  String get welcomeStoreTitle => 'Sovelluskeskuksen uusi ulkoasu';

  @override
  String get welcomeStoreSubtitle => 'Suosikkisovellusten löytäminen ei ole koskaan ollut yhtä helppoa.';

  @override
  String get welcomeSecurityTitle => 'Kohennettu tietoturva';

  @override
  String get welcomeSecuritySubtitle => 'Uusia asetuksia laitteistopohjaiselle salaukselle ja ZFS-salaukselle.';

  @override
  String welcomeChangelogLabel(Object url) {
    return '<a href=\"$url\">Näytä muutosloki</a>';
  }

  @override
  String get telemetryPageTitle => 'Telemetria';

  @override
  String telemetryHeader(Object distro) {
    return 'Auta parantamaan ${distro}a';
  }

  @override
  String telemetryDescription(Object distro) {
    return 'Jaa tietoja anonyymisti ${distro}n kanssa, jotta voimme parantaa kokemustasi.';
  }

  @override
  String telemetryLabelOn(Object distro) {
    return 'Kyllä, jaa järjestelmän tietoja $distro-tiimille';
  }

  @override
  String get telemetryLabelOff => 'Ei, älä jaa järjestelmän tietoja';

  @override
  String get telemetryReportLabel => 'Ilmoitettavat tiedot';

  @override
  String get telemetryReportTitle => 'Ilmoitettavat tiedot';

  @override
  String get telemetryLegalLabel => 'Lainopillinen';

  @override
  String get privacyPageTitle => 'Sijaintipalvelut';

  @override
  String get privacyLocationTitle => 'Otetaanko sijaintipalvelut käyttöön?';

  @override
  String get privacyLocationSubtitle => 'Salli sovellusten paikantaa sijaintisi.\nVoit muuttaa asetusta koska tahansa järjestelmän asetuksista.';

  @override
  String get privacyLocationEnable => 'Sijaintipalvelut';

  @override
  String get privacyPolicyLink => 'Tietosuoja';

  @override
  String ubuntuProMagicAttachInstructions(Object url) {
    return 'Liitä tämä erityisellä koodilla osoitteessa <a href=\"https://$url\">$url</a>';
  }

  @override
  String get ubuntuProPageTitle => 'Ubuntu Pro';

  @override
  String get ubuntuProHeader => 'Liitä tämä kone';

  @override
  String get ubuntuProOrAddTokenManually => 'tai lisää koodi (token) käsin';

  @override
  String get ubuntuProTokenTextfieldHint => 'Koodi';

  @override
  String get ubuntuProSubscriptionDescription => 'Hanki kokonaisvaltainen avoimen lähdekoodin tietoturvapalvelu';

  @override
  String ubuntuProSubscriptionFreeForPersonalUse(Object url) {
    return 'Henkilökohtaiseen käyttöön ilmainen, 30 päivän kokeilujakso yrityksille. <a href=\"https://$url\">Lue lisää</a>';
  }

  @override
  String get ubuntuProTokenAttachButton => 'Liitä';

  @override
  String get ubuntuProTokenAttachError => 'Virheellinen koodi';

  @override
  String get ubuntuProTokenAttachSucess => 'Virheellinen koodi';

  @override
  String get ubuntuProTokenAttachErrorVariant => 'Virheellinen koodi, yritä uudelleen';

  @override
  String get ubuntuProOnBoardingSkipForNow => 'Ohita toistaiseksi';

  @override
  String get ubuntuProOnBoardingSkipForNowDescription => 'Voit aina ottaa käyttöön Ubuntu Pron Ohjelmistot ja päivitykset -sovelluksessa.';

  @override
  String get ubuntuProOnBoardingEnableUbuntuPro => 'Ota Ubuntu Pro käyttöön';

  @override
  String get ubuntuProOnBoardingEnableUbuntuProDescription => 'Liitä tämä kone Ubuntu Pro -tilaukseen saadaksesi tietoturvapäivityksiä laajalle valikoimalle paketteja vuoteen 2032 asti.';

  @override
  String get ubuntuProSucessAttachHeader => 'Ubuntu Pro on nyt käytössä tälle tietokoneelle';

  @override
  String get ubuntuProSucessAttachDescription => 'Voit hallita Pro-palveluja Ohjelmistot ja päivitykset -sovelluksessa.';
}
