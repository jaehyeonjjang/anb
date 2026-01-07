import 'package:common_control/common_control.dart';
import 'package:periodic/models/apt.dart';

class Periodic {
  int id;
  String name;
  String aptname;
  String taskrange;
  String reportdate;
  String startdate;
  String enddate;
  String supply;
  String contract;
  String price;
  String safetygrade;
  int status;
  String prestartdate;
  String preenddate;
  String researchstartdate;
  String researchenddate;
  String analyzestartdate;
  String analyzeenddate;
  String ratingstartdate;
  String ratingenddate;
  String writestartdate;
  String writeenddate;
  String printstartdate;
  String printenddate;
  int blueprint1;
  int blueprint2;
  int blueprint3;
  int blueprint4;
  int blueprint5;
  int blueprint6;
  int blueprint7;
  int blueprint8;
  int blueprint9;
  String blueprint10;
  int blueprint11;
  String owner;
  String manager;
  String agent;
  int result1;
  int result2;
  int result3;
  int result4;
  int result5;
  String resulttext1;
  String resulttext2;
  String resulttext3;
  String resulttext4;
  String resulttext5;
  String past;
  Apt apt = Apt();
  String date;
  bool checked;
  Map<String, dynamic> extra;

  Periodic(
      {this.id = 0,
      this.name = '',
      this.aptname = '',
      this.taskrange = '',
      this.reportdate = '',
      this.startdate = '',
      this.enddate = '',
      this.supply = '',
      this.contract = '',
      this.price = '',
      this.safetygrade = '',
      this.status = 0,
      this.prestartdate = '',
      this.preenddate = '',
      this.researchstartdate = '',
      this.researchenddate = '',
      this.analyzestartdate = '',
      this.analyzeenddate = '',
      this.ratingstartdate = '',
      this.ratingenddate = '',
      this.writestartdate = '',
      this.writeenddate = '',
      this.printstartdate = '',
      this.printenddate = '',
      this.blueprint1 = 0,
      this.blueprint2 = 0,
      this.blueprint3 = 0,
      this.blueprint4 = 0,
      this.blueprint5 = 0,
      this.blueprint6 = 0,
      this.blueprint7 = 0,
      this.blueprint8 = 0,
      this.blueprint9 = 0,
      this.blueprint10 = '',
      this.blueprint11 = 0,
      this.owner = '',
      this.manager = '',
      this.agent = '',
      this.result1 = 0,
      this.result2 = 0,
      this.result3 = 0,
      this.result4 = 0,
      this.result5 = 0,
      this.resulttext1 = '',
      this.resulttext2 = '',
      this.resulttext3 = '',
      this.resulttext4 = '',
      this.resulttext5 = '',
      this.past = '',
      Apt? apt,
      this.date = '',
      this.extra = const {},
      this.checked = false}) {
    if (apt != null) {
      this.apt = apt;
    }
  }

  factory Periodic.fromJson(Map<String, dynamic> json) {
    Apt apt = Apt();

    if (json['extra'] != null && json['extra'] != '') {
      if (json['extra']['apt'] != null && json['extra']['apt'] != '') {
        apt = Apt.fromJson(json['extra']['apt']);
      }
    }

    return Periodic(
        id: json['id'] as int,
        name: json['name'] as String,
        aptname: json['aptname'] as String,
        taskrange: json['taskrange'] as String,
        reportdate: json['reportdate'] as String,
        startdate: json['startdate'] as String,
        enddate: json['enddate'] as String,
        supply: json['supply'] as String,
        contract: json['contract'] as String,
        price: json['price'] as String,
        safetygrade: json['safetygrade'] as String,
        status: json['status'] as int,
        prestartdate: json['prestartdate'] as String,
        preenddate: json['preenddate'] as String,
        researchstartdate: json['researchstartdate'] as String,
        researchenddate: json['researchenddate'] as String,
        analyzestartdate: json['analyzestartdate'] as String,
        analyzeenddate: json['analyzeenddate'] as String,
        ratingstartdate: json['ratingstartdate'] as String,
        ratingenddate: json['ratingenddate'] as String,
        writestartdate: json['writestartdate'] as String,
        writeenddate: json['writeenddate'] as String,
        printstartdate: json['printstartdate'] as String,
        printenddate: json['printenddate'] as String,
        blueprint1: json['blueprint1'] as int,
        blueprint2: json['blueprint2'] as int,
        blueprint3: json['blueprint3'] as int,
        blueprint4: json['blueprint4'] as int,
        blueprint5: json['blueprint5'] as int,
        blueprint6: json['blueprint6'] as int,
        blueprint7: json['blueprint7'] as int,
        blueprint8: json['blueprint8'] as int,
        blueprint9: json['blueprint9'] as int,
        blueprint10: json['blueprint10'] as String,
        blueprint11: json['blueprint11'] as int,
        owner: json['owner'] as String,
        manager: json['manager'] as String,
        agent: json['agent'] as String,
        result1: json['result1'] as int,
        result2: json['result2'] as int,
        result3: json['result3'] as int,
        result4: json['result4'] as int,
        result5: json['result5'] as int,
        resulttext1: json['resulttext1'] as String,
        resulttext2: json['resulttext2'] as String,
        resulttext3: json['resulttext3'] as String,
        resulttext4: json['resulttext4'] as String,
        resulttext5: json['resulttext5'] as String,
        past: json['past'] as String,
        apt: apt,
        date: json['date'] as String,
        extra: json['extra'] == null
            ? <String, dynamic>{}
            : json['extra'] as Map<String, dynamic>);
  }

  Map<String, dynamic> toJson() => {
        'id': id,
        'name': name,
        'aptname': aptname,
        'taskrange': taskrange,
        'reportdate': reportdate,
        'startdate': startdate,
        'enddate': enddate,
        'supply': supply,
        'contract': contract,
        'price': price,
        'safetygrade': safetygrade,
        'status': status,
        'prestartdate': prestartdate,
        'preenddate': preenddate,
        'researchstartdate': researchstartdate,
        'researchenddate': researchenddate,
        'analyzestartdate': analyzestartdate,
        'analyzeenddate': analyzeenddate,
        'ratingstartdate': ratingstartdate,
        'ratingenddate': ratingenddate,
        'writestartdate': writestartdate,
        'writeenddate': writeenddate,
        'printstartdate': printstartdate,
        'printenddate': printenddate,
        'blueprint1': blueprint1,
        'blueprint2': blueprint2,
        'blueprint3': blueprint3,
        'blueprint4': blueprint4,
        'blueprint5': blueprint5,
        'blueprint6': blueprint6,
        'blueprint7': blueprint7,
        'blueprint8': blueprint8,
        'blueprint9': blueprint9,
        'blueprint10': blueprint10,
        'blueprint11': blueprint11,
        'owner': owner,
        'manager': manager,
        'agent': agent,
        'result1': result1,
        'result2': result2,
        'result3': result3,
        'result4': result4,
        'result5': result5,
        'resulttext1': resulttext1,
        'resulttext2': resulttext2,
        'resulttext3': resulttext3,
        'resulttext4': resulttext4,
        'resulttext5': resulttext5,
        'past': past,
        'apt': apt.id,
        'date': date
      };
}

class PeriodicManager {
  static const baseUrl = '/api/periodic';

  static Future<List<Periodic>> find(
      {int page = 0, int pagesize = 20, String? params}) async {
    var result =
        await Http.get(baseUrl, {'page': page, 'pagesize': pagesize}, params);
    if (result == null || result['items'] == null) {
      return List<Periodic>.empty(growable: true);
    }

    return result['items']
        .map<Periodic>((json) => Periodic.fromJson(json))
        .toList();
  }

  static Future<Periodic> get(int id) async {
    var result = await Http.get('$baseUrl/$id');
    if (result == null || result['item'] == null) {
      return Periodic();
    }

    return Periodic.fromJson(result['item']);
  }

  static Future<int> insert(Periodic item) async {
    var result = await Http.insert(baseUrl, item.toJson());
    return result;
  }

  static update(Periodic item) async {
    await Http.put(baseUrl, item.toJson());
  }

  static delete(Periodic item) async {
    await Http.delete(baseUrl, item.toJson());
  }

  static Future<List<Periodic>> search(
      {int page = 0, int pagesize = 20, String? params}) async {
    var result = await Http.get(
        '$baseUrl/search', {'page': page, 'pagesize': pagesize}, params);
    if (result == null || result['items'] == null) {
      return List<Periodic>.empty(growable: true);
    }

    return result['items']
        .map<Periodic>((json) => Periodic.fromJson(json))
        .toList();
  }
}
