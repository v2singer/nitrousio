class GitProject {
  String id;
  String name;
  String url;
  String author;
  bool isSync;

  GitProject(
      {required this.id,
      required this.name,
      required this.url,
      required this.author,
      required this.isSync});
}
