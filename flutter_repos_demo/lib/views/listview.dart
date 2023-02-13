import 'package:flutter/material.dart';
import 'package:learn3/model/project.dart';

class ListViewBuilder extends StatelessWidget {
  const ListViewBuilder({
    super.key,
    required this.gitProjects,
  });

  final List<GitProject> gitProjects;

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
        itemCount: gitProjects.length,
        itemBuilder: (context, index) {
          return ListTile(title: Text(gitProjects[index].url));
        });
  }
}
