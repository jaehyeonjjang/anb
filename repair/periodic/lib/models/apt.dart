import 'package:common_control/common_control.dart';

class Apt {
  int id;
  String name;
  String completeyear;
  String flatcount;
  String type;
  String floor;
  String familycount;
  int familycount1;
  int familycount2;
  int familycount3;
  String tel;
  String fax;
  String email;
  String personalemail;
  String zip;
  String address;
  String address2;
  int contracttype;
  String contractprice;
  String testdate;
  String nexttestdate;
  String repair;
  String safety;
  String fault;
  String contractdate;
  String contractduration;
  String invoice;
  String depositdate;
  String fmsloginid;
  String fmspasswd;
  int facilitydivision;
  int facilitycategory;
  String position;
  String area;
  int groundfloor;
  int undergroundfloor;
  String useapproval;
  String date;
  bool checked;
  Map<String, dynamic> extra;

  Apt(
      {this.id = 0,
      this.name = '',
      this.completeyear = '',
      this.flatcount = '',
      this.type = '',
      this.floor = '',
      this.familycount = '',
      this.familycount1 = 0,
      this.familycount2 = 0,
      this.familycount3 = 0,
      this.tel = '',
      this.fax = '',
      this.email = '',
      this.personalemail = '',
      this.zip = '',
      this.address = '',
      this.address2 = '',
      this.contracttype = 0,
      this.contractprice = '',
      this.testdate = '',
      this.nexttestdate = '',
      this.repair = '',
      this.safety = '',
      this.fault = '',
      this.contractdate = '',
      this.contractduration = '',
      this.invoice = '',
      this.depositdate = '',
      this.fmsloginid = '',
      this.fmspasswd = '',
      this.facilitydivision = 0,
      this.facilitycategory = 0,
      this.position = '',
      this.area = '',
      this.groundfloor = 0,
      this.undergroundfloor = 0,
      this.useapproval = '',
      this.date = '',
      this.extra = const {},
      this.checked = false});

  factory Apt.fromJson(Map<String, dynamic> json) {
    return Apt(
        id: json['id'] as int,
        name: json['name'] as String,
        completeyear: json['completeyear'] as String,
        flatcount: json['flatcount'] as String,
        type: json['type'] as String,
        floor: json['floor'] as String,
        familycount: json['familycount'] as String,
        familycount1: json['familycount1'] as int,
        familycount2: json['familycount2'] as int,
        familycount3: json['familycount3'] as int,
        tel: json['tel'] as String,
        fax: json['fax'] as String,
        email: json['email'] as String,
        personalemail: json['personalemail'] as String,
        zip: json['zip'] as String,
        address: json['address'] as String,
        address2: json['address2'] as String,
        contracttype: json['contracttype'] as int,
        contractprice: json['contractprice'] as String,
        testdate: json['testdate'] as String,
        nexttestdate: json['nexttestdate'] as String,
        repair: json['repair'] as String,
        safety: json['safety'] as String,
        fault: json['fault'] as String,
        contractdate: json['contractdate'] as String,
        contractduration: json['contractduration'] as String,
        invoice: json['invoice'] as String,
        depositdate: json['depositdate'] as String,
        fmsloginid: json['fmsloginid'] as String,
        fmspasswd: json['fmspasswd'] as String,
        facilitydivision: json['facilitydivision'] as int,
        facilitycategory: json['facilitycategory'] as int,
        position: json['position'] as String,
        area: json['area'] as String,
        groundfloor: json['groundfloor'] as int,
        undergroundfloor: json['undergroundfloor'] as int,
        useapproval: json['useapproval'] as String,
        date: json['date'] as String,
        extra: json['extra'] == null ? <String, dynamic>{} : json['extra'] as Map<String, dynamic>);        
  }

  Map<String, dynamic> toJson() => {
        'id': id,
        'name': name,
        'completeyear': completeyear,
        'flatcount': flatcount,
        'type': type,
        'floor': floor,
        'familycount': familycount,
        'familycount1': familycount1,
        'familycount2': familycount2,
        'familycount3': familycount3,
        'tel': tel,
        'fax': fax,
        'email': email,
        'personalemail': personalemail,
        'zip': zip,
        'address': address,
        'address2': address2,
        'contracttype': contracttype,
        'contractprice': contractprice,
        'testdate': testdate,
        'nexttestdate': nexttestdate,
        'repair': repair,
        'safety': safety,
        'fault': fault,
        'contractdate': contractdate,
        'contractduration': contractduration,
        'invoice': invoice,
        'depositdate': depositdate,
        'fmsloginid': fmsloginid,
        'fmspasswd': fmspasswd,
        'facilitydivision': facilitydivision,
        'facilitycategory': facilitycategory,
        'position': position,
        'area': area,
        'groundfloor': groundfloor,
        'undergroundfloor': undergroundfloor,
        'useapproval': useapproval,
        'date': date
      };
}

class AptManager {
  static const baseUrl = '/api/apt';

  static Future<List<Apt>> find(
      {int page = 0, int pagesize = 20, String? params}) async {
    var result =
        await Http.get(baseUrl, {'page': page, 'pagesize': pagesize}, params);
    if (result == null || result['items'] == null) {
      return List<Apt>.empty(growable: true);
    }

    return result['items'].map<Apt>((json) => Apt.fromJson(json)).toList();
  }

  static Future<Apt> get(int id) async {
    var result = await Http.get('$baseUrl/$id');    
    if (result == null || result['item'] == null) {
      return Apt();
    }

    return Apt.fromJson(result['item']);
  }

  static Future<int> insert(Apt item) async {
    var result = await Http.insert(baseUrl, item.toJson());
    return result;
  }

  static update(Apt item) async {
    await Http.put(baseUrl, item.toJson());
  }

  static delete(Apt item) async {
    await Http.delete(baseUrl, item.toJson());
  }
}
